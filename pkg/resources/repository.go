// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Kerria

package resources

import (
	"fmt"
	"github.com/bmatcuk/doublestar/v4"
	krapi "github.com/kerria-dev/kerria/pkg/apis/kerria.dev"
	"github.com/kerria-dev/kerria/pkg/apis/kerria.dev/v1alpha1"
	"github.com/kerria-dev/kerria/pkg/openapi"
	"github.com/kerria-dev/kerria/pkg/util"
	"k8s.io/klog/v2"
	"os"
	"path/filepath"
	"reflect"
	kyyaml "sigs.k8s.io/kustomize/kyaml/yaml"
)

const (
	RepoFile = "kerria.repo.yaml"
	RepoKind = "Repository"
)

type Repository struct {
	Name           string
	GitRoot        string
	RepoRoot       string
	KustomizeFlags []string
	BuildPath      string
	Sources        []*Source
	Processors     []*Processor
}

type Source struct {
	ID          int
	Name        string
	Discoveries []*Discovery
}

type Discovery struct {
	Path        string
	Name        string
	Namespace   string
	Labels      map[string]string
	Annotations map[string]string
}

type Processor struct {
	ID            int
	Name          string
	Stage         ProcessorStage
	Properties    interface{}
	Image         string
	Network       bool
	StorageMounts []*StorageMount
	Env           []string
}

type ProcessorStage int

const (
	StageNone ProcessorStage = iota
	StagePreBuild
	StagePostBuild
)

var (
	ProcessorStages = map[string]ProcessorStage{
		"None":      StageNone,
		"PreBuild":  StagePreBuild,
		"PostBuild": StagePostBuild,
	}
	ProcessorStagesReverse map[ProcessorStage]string
)

type StorageMount struct {
	MountType     string
	Source        string
	Destination   string
	ReadWriteMode bool
}

const (
	fileKustomization = "kustomization.yaml"
	kindKustomization = "Kustomization"
	mountTypeGitRoot  = "bind"
	mountDestGitRoot  = "/gitroot"
)

// RepositoryFromAPI converts the latest API into the internal representation
func RepositoryFromAPI(apiRepo *v1alpha1.Repository) (*Repository, error) {
	repository := &Repository{}
	repository.Name = apiRepo.Name
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	repository.GitRoot, err = util.GitRepositoryRoot(cwd)
	if err != nil {
		return nil, err
	}
	repository.RepoRoot, err = filepath.Rel(repository.GitRoot, cwd)
	if err != nil {
		return nil, err
	}
	repository.KustomizeFlags = apiRepo.Spec.Build.KustomizeFlags
	repository.BuildPath = apiRepo.Spec.Build.OutputPath
	for idx, apiSourceConfig := range apiRepo.Spec.Sources {
		source := &Source{}
		source.ID = idx
		source.Name = apiSourceConfig.Name
		// v1alpha1 only supports glob patterns
		pattern := apiSourceConfig.Glob
		fsys := os.DirFS(".")
		globbed, err := doublestar.Glob(fsys, pattern,
			doublestar.WithFailOnIOErrors())
		if err != nil {
			return nil, err
		}
		for _, globDir := range globbed {
			rnode, err := kyyaml.ReadFile(filepath.Join(globDir, fileKustomization))
			if err != nil || rnode.GetKind() != kindKustomization {
				klog.Warningf("Skipped discovery for %s because no Kustomization was detected", globDir)
				continue
			}
			discovery := &Discovery{
				Path:        globDir,
				Name:        rnode.GetName(),
				Namespace:   rnode.GetNamespace(),
				Labels:      rnode.GetLabels(),
				Annotations: rnode.GetAnnotations(),
			}
			source.Discoveries = append(source.Discoveries, discovery)
		}
		repository.Sources = append(repository.Sources, source)
	}
	for idx, apiProcessor := range apiRepo.Spec.Processors {
		processor := &Processor{}
		processor.ID = idx
		processor.Name = apiProcessor.Name
		processor.Stage = ProcessorStages[string(apiProcessor.Stage)]
		processor.Properties = apiProcessor.Properties
		processor.Image = apiProcessor.Container.Image
		processor.Network = apiProcessor.Container.Network
		if apiProcessor.Container.MountGitRoot {
			processor.StorageMounts = append(processor.StorageMounts, &StorageMount{
				MountType:     mountTypeGitRoot,
				Source:        repository.GitRoot,
				Destination:   mountDestGitRoot,
				ReadWriteMode: true,
			})
		}
		for _, apiStorageMount := range apiProcessor.Container.AdditionalMounts {
			processor.StorageMounts = append(processor.StorageMounts, &StorageMount{
				MountType:     apiStorageMount.Type,
				Source:        apiStorageMount.Src,
				Destination:   apiStorageMount.Dst,
				ReadWriteMode: apiStorageMount.RW,
			})
		}
		processor.Env = apiProcessor.Container.Envs
		repository.Processors = append(repository.Processors, processor)
	}
	return repository, nil
}

func ReadRepository() (*Repository, error) {
	return ReadRepositoryWithPath(RepoFile)
}

func ReadRepositoryWithPath(path string) (repository *Repository, err error) {
	var rnode *kyyaml.RNode
	rnode, err = kyyaml.ReadFile(path)
	if err != nil {
		return
	}
	err = rnode.DeAnchor()
	if err != nil {
		return
	}
	apiVersion := rnode.GetApiVersion()
	kind := rnode.GetKind()
	if apiVersion != krapi.APIVersionV1Alpha1 {
		err = fmt.Errorf("unsupported apiVersion for Repository %s", apiVersion)
		return
	}
	if kind != RepoKind {
		err = fmt.Errorf("incorrect kind %s is not Repository", kind)
		return
	}
	typeMeta := kyyaml.TypeMeta{
		APIVersion: apiVersion,
		Kind:       kind,
	}
	var value reflect.Value
	value, err = openapi.DecodeValidatedDefaulted(rnode, typeMeta)
	if err != nil {
		return nil, err
	}
	v1alpha1Repo := value.Interface().(*v1alpha1.Repository)
	repository, err = RepositoryFromAPI(v1alpha1Repo)
	return
}

func init() {
	ProcessorStagesReverse = make(map[ProcessorStage]string, len(ProcessorStages))
	for key, value := range ProcessorStages {
		ProcessorStagesReverse[value] = key
	}
}
