// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Kerria

package build

import (
	"bytes"
	"embed"
	"github.com/Masterminds/sprig/v3"
	"github.com/kerria-dev/kerria/pkg/resources"
	"github.com/kerria-dev/kerria/pkg/scaffold"
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

func ReconcileDifferences(repository *resources.Repository, lockfile *resources.Lockfile, intersection *DiscoveryIntersection) ([]PairedDiscovery, error) {
	// Check discovery hashes for the resources that need rebuilds
	klog.Infof("Checking %v discovery directories...", len(intersection.Check))
	var keep []*resources.BuildStatus
	var needsRebuild []PairedDiscovery
	for _, pairing := range intersection.Check {
		digest, err := DirectoryHash(pairing.BuildStatus.DiscoHashType,
			repository.GitRoot, filepath.Join(repository.RepoRoot, pairing.BuildStatus.DiscoPath))
		if err != nil {
			return nil, err
		}
		if bytes.Equal(digest, pairing.BuildStatus.DiscoHash) {
			keep = append(keep, pairing.BuildStatus)
		} else {
			pairing.BuildStatus.DiscoHash = digest
			needsRebuild = append(needsRebuild, pairing)
		}
	}
	klog.Infof("%v discoveries need to be rebuilt", len(needsRebuild))

	// Remove all build directories associated with missing sources
	if len(intersection.Delete) > 0 {
		klog.Info("Removing builds with missing discoveries...")
	}
	err := removeBuilds(intersection.Delete)
	if err != nil {
		return nil, err
	}
	// Remove the build directories of what needs to be rebuilt
	if len(needsRebuild) > 0 {
		klog.Info("Removing builds that need to be rebuilt...")
	}
	err = removeBuilds(needsRebuild)
	if err != nil {
		return nil, err
	}

	// Build all requiring rebuild
	if len(needsRebuild) > 0 {
		klog.Info("Building what needs to be rebuilt...")
	}
	err = createBuilds(needsRebuild, repository.KustomizeFlags, repository.GitRoot, repository.RepoRoot)
	if err != nil {
		return nil, err
	}
	// Build all new discoveries
	if len(intersection.Create) > 0 {
		klog.Info("Building new discoveries...")
	}
	err = createBuilds(intersection.Create, repository.KustomizeFlags, repository.GitRoot, repository.RepoRoot)
	if err != nil {
		return nil, err
	}

	lockfile.Builds = []*resources.BuildStatus{}
	for _, pairing := range intersection.Check {
		lockfile.Builds = append(lockfile.Builds, pairing.BuildStatus)
	}
	for _, pairing := range intersection.Create {
		lockfile.Builds = append(lockfile.Builds, pairing.BuildStatus)
	}

	return needsRebuild, nil
}

func removeBuilds(pairs []PairedDiscovery) error {
	for _, pairing := range pairs {
		buildStatus := pairing.BuildStatus
		err := os.RemoveAll(buildStatus.BuildPath)
		if err != nil {
			return err
		}
		klog.Infof("Removed %s", buildStatus.BuildPath)
	}
	return nil
}

func createBuilds(pairs []PairedDiscovery, kustomizeFlags []string, repoRoot string, cwdRel string) error {
	for _, pairing := range pairs {
		buildStatus := pairing.BuildStatus
		klog.Infof("Building %s", buildStatus.DiscoPath)
		err := createBuild(buildStatus.DiscoPath, buildStatus.BuildPath, kustomizeFlags)
		if err != nil {
			return err
		}
		digest, err := DirectoryHash(buildStatus.BuildHashType,
			repoRoot, filepath.Join(cwdRel, buildStatus.BuildPath))
		if err != nil {
			return err
		}
		buildStatus.Timestamp = time.Now()
		buildStatus.BuildHash = digest
	}
	return nil
}

func createBuild(discoPath string, buildPath string, flags []string) error {
	output, err := KustomizeBuildCommand(discoPath, flags)
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
		DiscoPath string
	}{DiscoPath: discoPath})
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
