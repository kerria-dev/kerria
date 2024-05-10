// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Kerria

package openapi

import (
	"fmt"
	krapi "github.com/kerria-dev/kerria/pkg/apis/kerria.dev"
	"github.com/kerria-dev/kerria/pkg/apis/kerria.dev/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/kube-openapi/pkg/builder"
	"k8s.io/kube-openapi/pkg/builder3"
	"k8s.io/kube-openapi/pkg/common"
	"k8s.io/kube-openapi/pkg/spec3"
	"k8s.io/kube-openapi/pkg/validation/spec"
	"reflect"
	kyoapi "sigs.k8s.io/kustomize/kyaml/openapi"
	kyyaml "sigs.k8s.io/kustomize/kyaml/yaml"
	"strings"
)

const (
	extensionGVK             = "x-kubernetes-group-version-kind"
	internalAPIPrefix        = "github.com/kerria-dev/kerria/pkg/apis"
	kubernetesResourcePrefix = "io.k8s"
)

var (
	internalGroupVersionKinds map[string]v1.GroupVersionKind = nil
)

func ValueFromTypeMeta(typeMeta kyyaml.TypeMeta) (value reflect.Value, err error) {
	version, exists := krapi.Types[typeMeta.APIVersion]
	if !exists {
		err = fmt.Errorf("not a valid API version (%s) from TypeMeta", typeMeta.APIVersion)
		return
	}
	rtype, exists := version[typeMeta.Kind]
	if !exists {
		err = fmt.Errorf("not a valid Kind (%s) for API version (%s) from TypeMeta",
			typeMeta.Kind, typeMeta.APIVersion)
		return
	}
	value = reflect.New(rtype)
	return
}

func openAPIV3Scaffold() spec3.OpenAPI {
	return spec3.OpenAPI{
		Version: "3.0.0",
		Paths:   &spec3.Paths{},
		Info: &spec.Info{
			InfoProps: spec.InfoProps{
				Title:   "Kerria OpenAPI Specification",
				Version: "internal",
			},
		},
		Components: &spec3.Components{},
	}
}

func BuildInternalOpenAPIV3Specification() spec3.OpenAPI {
	api := openAPIV3Scaffold()
	api.Components.Schemas = BuildInternalOpenAPIV3Schemas()
	return api
}

func BuildInternalOpenAPIV3Schemas() map[string]*spec.Schema {
	config := common.OpenAPIV3Config{
		GetDefinitions:    GetOpenAPIDefinitions,
		GetDefinitionName: GetDefinitionName,
	}
	schemas, err := builder3.BuildOpenAPIDefinitionsForResources(&config, internalGVKNames()...)
	if err != nil {
		panic(err)
	}
	return schemas
}

func BuildInternalOpenAPIV2Specification() *spec.Swagger {
	config := common.Config{
		Info: &spec.Info{
			InfoProps: spec.InfoProps{
				Title:   "Kerria OpenAPI Specification",
				Version: "internal",
			},
		},
		GetDefinitions:    GetOpenAPIDefinitions,
		GetDefinitionName: GetDefinitionName,
	}

	definitions, err := builder.BuildOpenAPIDefinitionsForResources(&config, internalGVKNames()...)
	if err != nil {
		panic(err)
	}

	return definitions
}

func internalGVKNames() (internalGVKNames []string) {
	internalGVKNames = make([]string, 0, len(internalGroupVersionKinds))
	for _, gvk := range internalGroupVersionKinds {
		gvkName := fmt.Sprintf("%s/%s/%s.%s",
			internalAPIPrefix,
			gvk.Group, gvk.Version, gvk.Kind)
		internalGVKNames = append(internalGVKNames, gvkName)
	}
	return
}

func GetDefinitionName(name string) (string, spec.Extensions) {
	friendlyName := convertInternalName(name)
	if gvk, ok := internalGroupVersionKinds[friendlyName]; ok {
		exts := []interface{}{
			map[string]interface{}{
				"group":   gvk.Group,
				"version": gvk.Version,
				"kind":    gvk.Kind,
			},
		}

		return friendlyName, spec.Extensions{
			extensionGVK: exts,
		}
	}
	return friendlyName, nil
}

func convertInternalName(name string) string {
	parts := strings.Split(name, "/")
	if len(parts) > 0 {
		parts[0] = reverseDomain(parts[0])
		// If custom resource, extract G/V.K slice
		if !strings.HasPrefix(parts[0], kubernetesResourcePrefix) {
			parts = parts[len(parts)-2:]
			parts[0] = reverseDomain(parts[0])
		}
	}
	return strings.Join(parts, ".")
}

func reverseDomain(name string) string {
	if strings.Contains(name, ".") {
		parts := strings.Split(name, ".")
		for i, j := 0, len(parts)-1; i < j; i, j = i+1, j-1 {
			parts[i], parts[j] = parts[j], parts[i]
		}
		name = strings.Join(parts, ".")
	}
	return name
}

func ApplyGlobalSchema() {
	api := BuildInternalOpenAPIV2Specification()
	kyoapi.AddDefinitions(api.Definitions)
}

func init() {
	internalGroupVersionKinds = map[string]v1.GroupVersionKind{}
	for _, kind := range v1alpha1.Kinds {
		gvk := v1.GroupVersionKind{
			Group:   krapi.GroupName,
			Version: v1alpha1.Version,
			Kind:    kind,
		}
		internalGroupVersionKinds[fmt.Sprintf("%s.%s.%s", reverseDomain(gvk.Group), gvk.Version, gvk.Kind)] = gvk
	}
}
