// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Kerria

package meta

// TypeMeta describes an individual object in an API response or request with strings
// representing the type of the object and its API schema version.
type TypeMeta struct {
	// APIVersion defines the versioned schema of this representation of an object.
	APIVersion string `json:"apiVersion" yaml:"apiVersion"`
	// Kind is a string value representing the resource this object represents.
	Kind string `json:"kind" yaml:"kind"`
}

// ObjectMeta is metadata that all persisted resources must have, which includes all objects
// users must create.
type ObjectMeta struct {
	// Name is primarily intended for creation idempotence and configuration definition.
	Name string `json:"name", yaml:"name"`
	// Namespace defines the space within which each name must be unique.
	Namespace string `json:"namespace,omitempty" yaml:"namespace,omitempty"`
	// +default={}
	// Labels is a aap of string keys and values that can be used to organize and
	// categorize (scope and select) objects.
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	// +default={}
	// Annotations is an unstructured key value map stored with a resource that may
	// be set by external tools to store and retrieve arbitrary metadata.
	Annotations map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
}
