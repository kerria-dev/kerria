package openapi

import (
	"fmt"
	krapi "github.com/kerria-dev/kerria/pkg/apis/kerria.dev"
	"k8s.io/kube-openapi/pkg/validation/strfmt"
	"k8s.io/kube-openapi/pkg/validation/validate"
	"reflect"
	kyyaml "sigs.k8s.io/kustomize/kyaml/yaml"
)

func ValidateAgainstSchema(rnode *kyyaml.RNode, typeMeta kyyaml.TypeMeta) error {
	rschema := DereferenceResource(typeMeta)
	return validate.AgainstSchema(rschema.Schema, rnode, strfmt.NewFormats())
}

func SetDefaults(value reflect.Value, typeMeta kyyaml.TypeMeta) error {
	version, exists := krapi.Defaulters[typeMeta.APIVersion]
	if !exists {
		return fmt.Errorf("not a valid API version (%s) from TypeMeta", typeMeta.APIVersion)
	}
	defaulter, exists := version[typeMeta.Kind]
	if !exists {
		return fmt.Errorf("not a valid Kind (%s) for API version (%s) from TypeMeta",
			typeMeta.Kind, typeMeta.APIVersion)
	}
	defaulter.Call([]reflect.Value{value})
	return nil
}

func DecodeValidatedDefaulted(rnode *kyyaml.RNode, typeMeta kyyaml.TypeMeta) (value reflect.Value, err error) {
	err = ValidateAgainstSchema(rnode, typeMeta)
	if err != nil {
		return
	}
	value, err = ValueFromTypeMeta(typeMeta)
	if err != nil {
		return
	}
	err = rnode.YNode().Decode(value.Interface())
	if err != nil {
		return
	}
	err = SetDefaults(value, typeMeta)
	return
}
