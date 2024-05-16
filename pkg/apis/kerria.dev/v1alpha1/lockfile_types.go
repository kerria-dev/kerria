// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Kerria

package v1alpha1

import (
	"github.com/kerria-dev/kerria/pkg/apis/kerria.dev/meta"
)

// Lockfile is the current build state of the managed repository
type Lockfile struct {
	// +k8s:openapi-gen=true
	meta.TypeMeta `json:",inline" yaml:",inline"`
	// +k8s:openapi-gen=true
	meta.ObjectMeta `json:"metadata" yaml:"metadata"`

	Spec   LockfileSpec   `json:"spec,omitempty" yaml:"spec,omitempty"`
	Status LockfileStatus `json:"status,omitempty" yaml:"status,omitempty"`
}

// LockfileSpec represents the configuration for how to manage the lockfile
type LockfileSpec struct {
	// +default="sha256"
	DefaultHash string `json:"defaultHash,omitempty" yaml:"defaultHash,omitempty"`
}

// LockfileStatus represents the current state of the lockfile
type LockfileStatus struct {
	// +listType=map
	// +listMapKey=sourcePath
	// Builds is a list of BuildStatus for all builds in the repository
	Builds []BuildStatus `json:"builds" yaml:"builds"`
}

// BuildStatus defines the state of a build at the time it was built
type BuildStatus struct {
	// Timestamp is the RFC 3339 timestamp for when the build was performed
	Timestamp string `json:"timestamp" yaml:"timestamp"`
	// SourceHash is the source directory hash digest in the form of SourceHashType
	SourceHash string `json:"sourceHash" yaml:"sourceHash"`
	// SourceHashType is the type of hash used to compute the source digest
	SourceHashType HashAlgorithm `json:"sourceHashType" yaml:"sourceHashType"`
	// SourcePath is the path to the source directory used to create the build
	SourcePath string `json:"sourcePath" yaml:"sourcePath"`
	// BuildHash is the build directory hash digest in the form of BuildHashType
	BuildHash string `json:"buildHash" yaml:"buildHash"`
	// BuildHashType is the type of hash used to compute the build digest
	BuildHashType HashAlgorithm `json:"buildHashType" yaml:"buildHashType"`
	// BuildPath is the path to the build directory where the build output was written to
	BuildPath string `json:"buildPath" yaml:"buildPath"`
}

// HashAlgorithm is a string enum representing the supported hash algorithms
// +enum
type HashAlgorithm string

const (
	MD5    HashAlgorithm = "md5"
	SHA1   HashAlgorithm = "sha1"
	SHA256 HashAlgorithm = "sha256"
	SHA512 HashAlgorithm = "sha512"
)
