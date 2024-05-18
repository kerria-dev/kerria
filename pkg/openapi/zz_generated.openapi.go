//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Kerria

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package openapi

import (
	common "k8s.io/kube-openapi/pkg/common"
	spec "k8s.io/kube-openapi/pkg/validation/spec"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/kerria-dev/kerria/pkg/apis/kerria.dev/meta.ObjectMeta":                 schema_pkg_apis_kerriadev_meta_ObjectMeta(ref),
		"github.com/kerria-dev/kerria/pkg/apis/kerria.dev/meta.TypeMeta":                   schema_pkg_apis_kerriadev_meta_TypeMeta(ref),
		"github.com/kerria-dev/kerria/pkg/apis/kerria.dev/v1alpha1.BuildStatus":            schema_pkg_apis_kerriadev_v1alpha1_BuildStatus(ref),
		"github.com/kerria-dev/kerria/pkg/apis/kerria.dev/v1alpha1.ContainerSpec":          schema_pkg_apis_kerriadev_v1alpha1_ContainerSpec(ref),
		"github.com/kerria-dev/kerria/pkg/apis/kerria.dev/v1alpha1.Lockfile":               schema_pkg_apis_kerriadev_v1alpha1_Lockfile(ref),
		"github.com/kerria-dev/kerria/pkg/apis/kerria.dev/v1alpha1.LockfileSpec":           schema_pkg_apis_kerriadev_v1alpha1_LockfileSpec(ref),
		"github.com/kerria-dev/kerria/pkg/apis/kerria.dev/v1alpha1.LockfileStatus":         schema_pkg_apis_kerriadev_v1alpha1_LockfileStatus(ref),
		"github.com/kerria-dev/kerria/pkg/apis/kerria.dev/v1alpha1.ProcessorConfig":        schema_pkg_apis_kerriadev_v1alpha1_ProcessorConfig(ref),
		"github.com/kerria-dev/kerria/pkg/apis/kerria.dev/v1alpha1.Repository":             schema_pkg_apis_kerriadev_v1alpha1_Repository(ref),
		"github.com/kerria-dev/kerria/pkg/apis/kerria.dev/v1alpha1.RepositoryBuildConfig":  schema_pkg_apis_kerriadev_v1alpha1_RepositoryBuildConfig(ref),
		"github.com/kerria-dev/kerria/pkg/apis/kerria.dev/v1alpha1.RepositorySourceConfig": schema_pkg_apis_kerriadev_v1alpha1_RepositorySourceConfig(ref),
		"github.com/kerria-dev/kerria/pkg/apis/kerria.dev/v1alpha1.RepositorySpec":         schema_pkg_apis_kerriadev_v1alpha1_RepositorySpec(ref),
		"github.com/kerria-dev/kerria/pkg/apis/kerria.dev/v1alpha1.StorageMount":           schema_pkg_apis_kerriadev_v1alpha1_StorageMount(ref),
	}
}

func schema_pkg_apis_kerriadev_meta_ObjectMeta(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"name": {
						SchemaProps: spec.SchemaProps{
							Description: "Name is primarily intended for creation idempotence and configuration definition.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"namespace": {
						SchemaProps: spec.SchemaProps{
							Description: "Namespace defines the space within which each name must be unique.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"labels": {
						SchemaProps: spec.SchemaProps{
							Description: "Labels is a aap of string keys and values that can be used to organize and categorize (scope and select) objects.",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
					"annotations": {
						SchemaProps: spec.SchemaProps{
							Description: "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
				},
				Required: []string{"name"},
			},
		},
	}
}

func schema_pkg_apis_kerriadev_meta_TypeMeta(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "TypeMeta describes an individual object in an API response or request with strings representing the type of the object and its API schema version.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the resource this object represents.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"apiVersion", "kind"},
			},
		},
	}
}

func schema_pkg_apis_kerriadev_v1alpha1_BuildStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "BuildStatus defines the state of a build at the time it was built",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"timestamp": {
						SchemaProps: spec.SchemaProps{
							Description: "Timestamp is the RFC 3339 timestamp for when the build was performed",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"sourceHash": {
						SchemaProps: spec.SchemaProps{
							Description: "SourceHash is the source directory hash digest in the form of SourceHashType",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"sourceHashType": {
						SchemaProps: spec.SchemaProps{
							Description: "SourceHashType is the type of hash used to compute the source digest\n\nPossible enum values:\n - `\"md5\"`\n - `\"sha1\"`\n - `\"sha256\"`\n - `\"sha512\"`",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
							Enum:        []interface{}{"md5", "sha1", "sha256", "sha512"},
						},
					},
					"sourcePath": {
						SchemaProps: spec.SchemaProps{
							Description: "SourcePath is the path to the source directory used to create the build",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"buildHash": {
						SchemaProps: spec.SchemaProps{
							Description: "BuildHash is the build directory hash digest in the form of BuildHashType",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"buildHashType": {
						SchemaProps: spec.SchemaProps{
							Description: "BuildHashType is the type of hash used to compute the build digest\n\nPossible enum values:\n - `\"md5\"`\n - `\"sha1\"`\n - `\"sha256\"`\n - `\"sha512\"`",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
							Enum:        []interface{}{"md5", "sha1", "sha256", "sha512"},
						},
					},
					"buildPath": {
						SchemaProps: spec.SchemaProps{
							Description: "BuildPath is the path to the build directory where the build output was written to",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"timestamp", "sourceHash", "sourceHashType", "sourcePath", "buildHash", "buildHashType", "buildPath"},
			},
		},
	}
}

func schema_pkg_apis_kerriadev_v1alpha1_ContainerSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ContainerSpec defines a spec for running a function as a container",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"image": {
						SchemaProps: spec.SchemaProps{
							Description: "Image is the container image to run",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"network": {
						SchemaProps: spec.SchemaProps{
							Description: "Network defines network specific configuration",
							Default:     false,
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"mountGitRoot": {
						SchemaProps: spec.SchemaProps{
							Description: "MountGitRoot defines whether Kerria will automatically mount the directory of repo file into the processor. The destination for this automatic mount is /gitroot within the container",
							Default:     true,
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"additionalMounts": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-map-keys": []interface{}{
									"src",
								},
								"x-kubernetes-list-type": "map",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "AdditionalMounts are additional storage or directories to mount into the container",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("github.com/kerria-dev/kerria/pkg/apis/kerria.dev/v1alpha1.StorageMount"),
									},
								},
							},
						},
					},
					"envs": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "atomic",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "Envs is a slice of env string that will be exposed to container",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
				},
				Required: []string{"image"},
			},
		},
		Dependencies: []string{
			"github.com/kerria-dev/kerria/pkg/apis/kerria.dev/v1alpha1.StorageMount"},
	}
}

func schema_pkg_apis_kerriadev_v1alpha1_Lockfile(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Lockfile is the current build state of the managed repository",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the resource this object represents.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("github.com/kerria-dev/kerria/pkg/apis/kerria.dev/meta.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("github.com/kerria-dev/kerria/pkg/apis/kerria.dev/v1alpha1.LockfileSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("github.com/kerria-dev/kerria/pkg/apis/kerria.dev/v1alpha1.LockfileStatus"),
						},
					},
				},
				Required: []string{"apiVersion", "kind", "metadata"},
			},
		},
		Dependencies: []string{
			"github.com/kerria-dev/kerria/pkg/apis/kerria.dev/meta.ObjectMeta", "github.com/kerria-dev/kerria/pkg/apis/kerria.dev/v1alpha1.LockfileSpec", "github.com/kerria-dev/kerria/pkg/apis/kerria.dev/v1alpha1.LockfileStatus"},
	}
}

func schema_pkg_apis_kerriadev_v1alpha1_LockfileSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "LockfileSpec represents the configuration for how to manage the lockfile",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"defaultHash": {
						SchemaProps: spec.SchemaProps{
							Default: "sha256",
							Type:    []string{"string"},
							Format:  "",
						},
					},
				},
			},
		},
	}
}

func schema_pkg_apis_kerriadev_v1alpha1_LockfileStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "LockfileStatus represents the current state of the lockfile",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"builds": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-map-keys": []interface{}{
									"sourcePath",
								},
								"x-kubernetes-list-type": "map",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "Builds is a list of BuildStatus for all builds in the repository",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("github.com/kerria-dev/kerria/pkg/apis/kerria.dev/v1alpha1.BuildStatus"),
									},
								},
							},
						},
					},
				},
				Required: []string{"builds"},
			},
		},
		Dependencies: []string{
			"github.com/kerria-dev/kerria/pkg/apis/kerria.dev/v1alpha1.BuildStatus"},
	}
}

func schema_pkg_apis_kerriadev_v1alpha1_ProcessorConfig(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ProcessorConfig defines how a processor should be configured for the repository",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"name": {
						SchemaProps: spec.SchemaProps{
							Default: "",
							Type:    []string{"string"},
							Format:  "",
						},
					},
					"stage": {
						SchemaProps: spec.SchemaProps{
							Description: "Stage is the build stage to run this processor in\n\nPossible enum values:\n - `\"None\"`\n - `\"PostBuild\"`\n - `\"PreBuild\"`",
							Default:     "None",
							Type:        []string{"string"},
							Format:      "",
							Enum:        []interface{}{"None", "PostBuild", "PreBuild"},
						},
					},
					"container": {
						SchemaProps: spec.SchemaProps{
							Description: "Container defines the container that executes the processor",
							Default:     map[string]interface{}{},
							Ref:         ref("github.com/kerria-dev/kerria/pkg/apis/kerria.dev/v1alpha1.ContainerSpec"),
						},
					},
					"properties": {
						SchemaProps: spec.SchemaProps{
							Description: "Properties are arbitrary properties passed directly to the processor",
							Type:        []string{"object"},
							Format:      "",
						},
					},
				},
				Required: []string{"name", "container"},
			},
		},
		Dependencies: []string{
			"github.com/kerria-dev/kerria/pkg/apis/kerria.dev/v1alpha1.ContainerSpec"},
	}
}

func schema_pkg_apis_kerriadev_v1alpha1_Repository(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Repository is the configuration API for a kerria-managed repository",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the resource this object represents.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("github.com/kerria-dev/kerria/pkg/apis/kerria.dev/meta.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("github.com/kerria-dev/kerria/pkg/apis/kerria.dev/v1alpha1.RepositorySpec"),
						},
					},
				},
				Required: []string{"apiVersion", "kind", "metadata"},
			},
		},
		Dependencies: []string{
			"github.com/kerria-dev/kerria/pkg/apis/kerria.dev/meta.ObjectMeta", "github.com/kerria-dev/kerria/pkg/apis/kerria.dev/v1alpha1.RepositorySpec"},
	}
}

func schema_pkg_apis_kerriadev_v1alpha1_RepositoryBuildConfig(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "RepositoryBuildConfig specifies how kerria should build and store manifests",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kustomizeFlags": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "atomic",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "KustomizeFlags are a list of additional flags to add to the `kustomize build` command",
							Default:     []interface{}{"--enable-helm", "--enable-alpha-plugins", "--network"},
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
					"outputPath": {
						SchemaProps: spec.SchemaProps{
							Description: "OutputPath is the directory kerria uses as the root of the build output tree",
							Default:     "builds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
			},
		},
	}
}

func schema_pkg_apis_kerriadev_v1alpha1_RepositorySourceConfig(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "RepositorySourceConfig describes how kerria should find source manifests",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"name": {
						SchemaProps: spec.SchemaProps{
							Description: "Name is an optional name for the source",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"glob": {
						SchemaProps: spec.SchemaProps{
							Description: "Glob is a directory pattern for manifest sources",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"glob"},
			},
		},
	}
}

func schema_pkg_apis_kerriadev_v1alpha1_RepositorySpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "RepositorySpec is the specification for how kerria manages a repository",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"build": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("github.com/kerria-dev/kerria/pkg/apis/kerria.dev/v1alpha1.RepositoryBuildConfig"),
						},
					},
					"sources": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-map-keys": []interface{}{
									"path",
								},
								"x-kubernetes-list-type": "map",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "Sources is a list of RepositorySourceConfig to find source manifests",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("github.com/kerria-dev/kerria/pkg/apis/kerria.dev/v1alpha1.RepositorySourceConfig"),
									},
								},
							},
						},
					},
					"processors": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-map-keys": []interface{}{
									"name",
								},
								"x-kubernetes-list-type": "map",
							},
						},
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("github.com/kerria-dev/kerria/pkg/apis/kerria.dev/v1alpha1.ProcessorConfig"),
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/kerria-dev/kerria/pkg/apis/kerria.dev/v1alpha1.ProcessorConfig", "github.com/kerria-dev/kerria/pkg/apis/kerria.dev/v1alpha1.RepositoryBuildConfig", "github.com/kerria-dev/kerria/pkg/apis/kerria.dev/v1alpha1.RepositorySourceConfig"},
	}
}

func schema_pkg_apis_kerriadev_v1alpha1_StorageMount(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "StorageMount represents a container's mounted storage option(s)",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"type": {
						SchemaProps: spec.SchemaProps{
							Description: "Type of mount e.g. bind mount, local volume, etc.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"src": {
						SchemaProps: spec.SchemaProps{
							Description: "Src for the storage to be mounted. For named volumes, this is the name of the volume. For anonymous volumes, this field is omitted (empty string). For bind mounts, this is the path to the file or directory on the host.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"dst": {
						SchemaProps: spec.SchemaProps{
							Description: "Dst where the file or directory is mounted in the container.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"rw": {
						SchemaProps: spec.SchemaProps{
							Description: "RW to mount in ReadWrite mode if it's explicitly configured See https://docs.docker.com/storage/bind-mounts/#use-a-read-only-bind-mount",
							Default:     false,
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
				},
			},
		},
	}
}
