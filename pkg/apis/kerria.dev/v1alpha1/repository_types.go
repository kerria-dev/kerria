// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Kerria

package v1alpha1

import (
	"github.com/kerria-dev/kerria/pkg/apis/kerria.dev/meta"
)

// Repository is the configuration API for a kerria-managed repository
type Repository struct {
	meta.TypeMeta   `json:",inline" yaml:",inline"`
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
	// +listType=map
	// +listMapKey=name
	Processors []ProcessorConfig `json:"processors,omitempty" yaml:"processors,omitempty"`
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

// ProcessorConfig defines how a processor should be configured for the repository
type ProcessorConfig struct {
	Name string `json:"name" yaml:"name"`
	// +default="None"
	// Stage is the build stage to run this processor in
	Stage ProcessorStage `json:"stage,omitempty" yaml:"stage,omitempty"`
	// Container defines the container that executes the processor
	Container ContainerSpec `json:"container" yaml:"container"`
	// Properties are arbitrary properties passed directly to the processor
	Properties interface{} `json:"properties,omitempty" yaml:"properties,omitempty"`
}

// ProcessorStage is a string enum representing the build process stage a processor should run in
// +enum
type ProcessorStage string

const (
	None      ProcessorStage = "None"
	PreBuild  ProcessorStage = "PreBuild"
	PostBuild ProcessorStage = "PostBuild"
)

// ContainerSpec defines a spec for running a function as a container
type ContainerSpec struct {
	// Image is the container image to run
	Image string `json:"image" yaml:"image"`

	// +default=false
	// Network defines network specific configuration
	Network bool `json:"network,omitempty" yaml:"network,omitempty"`

	// +default=true
	// MountGitRoot defines whether Kerria will automatically mount the directory of repo file into the processor.
	// The destination for this automatic mount is /gitroot within the container
	MountGitRoot bool `json:"mountGitRoot,omitempty" yaml:"mountGitRoot,omitempty"`

	// +listType=map
	// +listMapKey=src
	// AdditionalMounts are additional storage or directories to mount into the container
	AdditionalMounts []StorageMount `json:"additionalMounts,omitempty" yaml:"additionalMounts,omitempty"`

	// +listType=atomic
	// Envs is a slice of env string that will be exposed to container
	Envs []string `json:"envs,omitempty" yaml:"envs,omitempty"`
}

// StorageMount represents a container's mounted storage option(s)
type StorageMount struct {
	// Type of mount e.g. bind mount, local volume, etc.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// Src for the storage to be mounted.
	// For named volumes, this is the name of the volume.
	// For anonymous volumes, this field is omitted (empty string).
	// For bind mounts, this is the path to the file or directory on the host.
	Src string `json:"src,omitempty" yaml:"src,omitempty"`

	// Dst where the file or directory is mounted in the container.
	Dst string `json:"dst,omitempty" yaml:"dst,omitempty"`

	// +default=false
	// RW to mount in ReadWrite mode if it's explicitly configured
	// See https://docs.docker.com/storage/bind-mounts/#use-a-read-only-bind-mount
	RW bool `json:"rw,omitempty" yaml:"rw,omitempty"`
}
