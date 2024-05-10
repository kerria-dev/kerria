// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Kerria

package openapi

import (
	"k8s.io/kube-openapi/pkg/validation/spec"
	kyoapi "sigs.k8s.io/kustomize/kyaml/openapi"
	kyyaml "sigs.k8s.io/kustomize/kyaml/yaml"
)

const (
	extensionDeref = "x-kerria-dereferenced-from"
)

func DereferenceResource(rtype kyyaml.TypeMeta) *kyoapi.ResourceSchema {
	return DereferenceResourceWithDepth(rtype, 0)
}

func DereferenceResourceWithDepth(resource kyyaml.TypeMeta, depth int) *kyoapi.ResourceSchema {
	iterations := 0

	root := kyoapi.SchemaForResourceType(resource)
	if root == nil {
		return nil
	}

	stack := []*kyoapi.ResourceSchema{root}

	for len(stack) > 0 && // While stack is not empty
		((depth <= 0) || (iterations < depth)) { // Supports both depth limit and disable (0)
		// Pop
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// Process and Push
		for key, value := range current.Schema.SchemaProps.Properties {
			// For map-like schemata
			field := current.Lookup(key)
			if field != nil {
				replacement := dereferenceSchema(field, &value)
				if replacement != nil {
					current.Schema.SchemaProps.Properties[key] = *replacement.Schema
					stack = append(stack, replacement)
				}
			}

			// For array-like schemata
			elements := current.Lookup(key, "[]")
			if elements != nil {
				replacement := dereferenceSchema(elements, value.Items.Schema)
				if replacement != nil {
					value.Items.Schema = replacement.Schema
					stack = append(stack, replacement)
				}
			}
		}

		iterations++
	}

	return root
}

func dereferenceSchema(schema *kyoapi.ResourceSchema, propertyValue *spec.Schema) (replacement *kyoapi.ResourceSchema) {
	ref := &propertyValue.SchemaProps.Ref
	if ref.GetPointer().IsEmpty() {
		return
	}

	if schema.Schema.Extensions == nil {
		schema.Schema.Extensions = make(spec.Extensions)
	}
	schema.Schema.Extensions.Add(extensionDeref, ref.String())
	replacement = schema

	return
}
