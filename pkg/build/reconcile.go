// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Kerria

package build

import (
	"bytes"
	"embed"
	"github.com/Masterminds/sprig/v3"
	"github.com/kerria-dev/kerria/pkg/resources"
	"github.com/kerria-dev/kerria/pkg/scaffold"
	"github.com/kerria-dev/kerria/pkg/util"
	"k8s.io/klog/v2"
	"os"
	"path/filepath"
	"text/template"
	"time"
)

var (
	//go:embed templates/README.md.tmpl
	templateReadmeEmbed embed.FS
	templateReadme      *template.Template
)

func ReconcileDifferences(repository *resources.Repository, lockfile *resources.Lockfile, diff *LockfileDiff) error {
	klog.Info("Reconciling...")
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	repoRoot, err := util.GitRepositoryRoot(cwd)
	if err != nil {
		return err
	}
	cwdRel, err := filepath.Rel(repoRoot, cwd)
	if err != nil {
		return err
	}

	// Check source hashes for the resources that need rebuilds
	klog.Infof("Checking %v source directories...", len(diff.Check))
	var keep []*resources.BuildStatus
	var needsRebuild []*resources.BuildStatus
	for _, buildStatus := range diff.Check {
		digest, err := util.DirectoryHash(buildStatus.SourceHashType,
			repoRoot, filepath.Join(cwdRel, buildStatus.SourcePath))
		if err != nil {
			return err
		}
		if bytes.Equal(digest, buildStatus.SourceHash) {
			keep = append(keep, buildStatus)
		} else {
			buildStatus.SourceHash = digest
			needsRebuild = append(needsRebuild, buildStatus)
		}
	}
	klog.Infof("%v sources need to be rebuilt", len(needsRebuild))

	// Remove all build directories associated with missing sources
	klog.Info("Removing builds with missing sources...")
	err = removeBuilds(diff.Delete)
	if err != nil {
		return err
	}
	// Remove the build directories of what needs to be rebuilt
	klog.Info("Removing builds that need to be rebuilt...")
	err = removeBuilds(needsRebuild)
	if err != nil {
		return err
	}

	// Build all requiring rebuild
	klog.Info("Building what needs to be rebuilt...")
	err = createBuilds(needsRebuild, repository.KustomizeFlags, repoRoot, cwdRel)
	if err != nil {
		return err
	}
	// Build all new sources
	klog.Info("Building new sources...")
	err = createBuilds(diff.Create, repository.KustomizeFlags, repoRoot, cwdRel)
	if err != nil {
		return err
	}

	lockfile.Builds = append(diff.Check, diff.Create...)

	return nil
}

func removeBuilds(buildStatuses []*resources.BuildStatus) error {
	for _, buildStatus := range buildStatuses {
		err := os.RemoveAll(buildStatus.BuildPath)
		if err != nil {
			return err
		}
		klog.Infof("Removed %s", buildStatus.BuildPath)
	}
	return nil
}

func createBuilds(buildStatuses []*resources.BuildStatus, kustomizeFlags []string, repoRoot string, cwdRel string) error {
	for _, buildStatus := range buildStatuses {
		klog.Infof("Building %s", buildStatus.SourcePath)
		err := createBuild(buildStatus.SourcePath, buildStatus.BuildPath, kustomizeFlags)
		if err != nil {
			return err
		}
		digest, err := util.DirectoryHash(buildStatus.BuildHashType,
			repoRoot, filepath.Join(cwdRel, buildStatus.SourcePath))
		if err != nil {
			return err
		}
		buildStatus.Timestamp = time.Now()
		buildStatus.BuildHash = digest
	}
	return nil
}

func createBuild(sourcePath string, buildPath string, flags []string) error {
	output, err := KustomizeBuildCommand(sourcePath, flags)
	if err != nil {
		return err
	}
	err = os.MkdirAll(buildPath, 0755)
	if err != nil {
		return err
	}
	fileReadme, err := os.Create(filepath.Join(buildPath, "README.md"))
	if err != nil {
		return err
	}
	defer fileReadme.Close()
	err = templateReadme.Execute(fileReadme, struct {
		SourcePath string
	}{SourcePath: sourcePath})
	if err != nil {
		return err
	}
	fileBuild, err := os.Create(filepath.Join(buildPath, "build.yaml"))
	if err != nil {
		return err
	}
	defer fileBuild.Close()
	_, err = fileBuild.WriteString(output)
	if err != nil {
		return err
	}
	kustomization, err := scaffold.Encode(scaffold.ScaffoldKustomization())
	if err != nil {
		return err
	}
	fileKustomization, err := os.Create(filepath.Join(buildPath, "kustomization.yaml"))
	if err != nil {
		return err
	}
	defer fileKustomization.Close()
	_, err = fileKustomization.Write(kustomization)
	if err != nil {
		return err
	}
	return nil
}

func init() {
	templateReadmeContent, err := templateReadmeEmbed.ReadFile("templates/README.md.tmpl")
	if err != nil {
		panic(err)
	}
	templateReadme, err = template.New("build-README").Funcs(sprig.FuncMap()).Parse(string(templateReadmeContent))
	if err != nil {
		panic(err)
	}
}
