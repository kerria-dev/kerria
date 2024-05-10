// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Kerria

package resources

import (
	"fmt"
	"github.com/bmatcuk/doublestar/v4"
	krapi "github.com/kerria-dev/kerria/pkg/apis/kerria.dev"
	"github.com/kerria-dev/kerria/pkg/apis/kerria.dev/v1alpha1"
	"github.com/kerria-dev/kerria/pkg/openapi"
	"os"
	"reflect"
	kyyaml "sigs.k8s.io/kustomize/kyaml/yaml"
)

const (
	RepoFile = "kerria.repo.yaml"
	RepoKind = "Repository"
)

type Repository struct {
	Name           string
	KustomizeFlags []string
	BuildPath      string
	Sources        []*Source
}

type Source struct {
	ID    int
	Name  string
	Paths []string
}

// RepositoryFromAPI converts the latest API into the internal representation
func RepositoryFromAPI(apiRepo *v1alpha1.Repository) (repository *Repository, err error) {
	repository = &Repository{}
	repository.Name = apiRepo.Name
	repository.KustomizeFlags = apiRepo.Spec.Build.KustomizeFlags
	repository.BuildPath = apiRepo.Spec.Build.OutputPath
	for idx, apiSourceConfig := range apiRepo.Spec.Sources {
		source := Source{}
		source.ID = idx
		source.Name = apiSourceConfig.Name
		// v1alpha1 only supports glob patterns
		pattern := apiSourceConfig.Glob

		fsys := os.DirFS(".")
		source.Paths, err = doublestar.Glob(fsys, pattern,
			doublestar.WithFailOnIOErrors())
		if err != nil {
			return
		}

		repository.Sources = append(repository.Sources, &source)
	}
	return
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
