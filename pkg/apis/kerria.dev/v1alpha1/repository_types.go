// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Kerria

package v1alpha1

import (
	"github.com/kerria-dev/kerria/pkg/apis/kerria.dev/meta"
)

// Repository is the configuration API for a kerria-managed repository
type Repository struct {
	// +k8s:openapi-gen=true
	meta.TypeMeta `json:",inline" yaml:",inline"`
	// +k8s:openapi-gen=true
	meta.ObjectMeta `json:"metadata" yaml:"metadata"`

	Spec RepositorySpec `json:"spec,omitempty" yaml:"spec,omitempty"`
}

// RepositorySpec is the specification for how kerria manages a repository
type RepositorySpec struct {
	Build RepositoryBuildConfig `json:"build,omitempty" yaml:"build,omitempty"`
	// +listType=map
	// +listMapKey=path
	// Sources is a list of RepositorySourceConfig to find source manifests
	Sources []RepositorySourceConfig `json:"sources,omitempty" yaml:"sources,omitempty"`
}

// RepositoryBuildConfig specifies how kerria should build and store manifests
type RepositoryBuildConfig struct {
	// +listType=atomic
	// +default=["--enable-helm", "--enable-alpha-plugins", "--network"]
	// KustomizeFlags are a list of additional flags to add to the `kustomize build` command
	KustomizeFlags []string `json:"kustomizeFlags,omitempty" yaml:"kustomizeFlags,omitempty"`

	// +default="builds"
	// OutputPath is the directory kerria uses as the root of the build output tree
	OutputPath string `json:"outputPath,omitempty" yaml:"outputPath,omitempty"`
}

// RepositorySourceConfig describes how kerria should find source manifests
type RepositorySourceConfig struct {
	// Name is an optional name for the source
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	// Glob is a directory pattern for manifest sources
	Glob string `json:"glob" yaml:"glob"`
}
