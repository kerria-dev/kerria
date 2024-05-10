// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Kerria

package cmd

import (
	"embed"
	"github.com/Masterminds/sprig/v3"
	"github.com/kerria-dev/kerria/cmd/util"
	"github.com/kerria-dev/kerria/pkg/openapi"
	"github.com/spf13/cobra"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/kube-openapi/pkg/validation/spec"
	"os"
	kyyaml "sigs.k8s.io/kustomize/kyaml/yaml"
	"text/template"
)

//go:embed templates/explain.tmpl
var templateExplain embed.FS

var explainCmd = &cobra.Command{
	Use:   "explain",
	Short: "Display API documentation about Kerria resources",
	Run: func(cmd *cobra.Command, args []string) {
		rschema := openapi.DereferenceResourceWithDepth(kyyaml.TypeMeta{
			APIVersion: "kerria.dev/v1alpha1",
			Kind:       "Repository",
		}, 1)
		if rschema == nil {
			panic("rschema is nil")
		}

		templateContent, err := templateExplain.ReadFile("templates/explain.tmpl")
		if err != nil {
			panic(err)
		}
		template, err := template.New("explain").Funcs(sprig.FuncMap()).Parse(string(templateContent))
		if err != nil {
			panic(err)
		}

		err = template.Execute(os.Stdout, explainData{
			GVK: v1.GroupVersionKind{
				Group:   "kerria.dev",
				Version: "v1alpha1",
				Kind:    "Repository",
			},
			Schema: *rschema.Schema,
			Colors: colors,
		})
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(explainCmd)
}

type explainData struct {
	GVK    v1.GroupVersionKind
	Schema spec.Schema
	Field  *fieldData
	Colors util.ANSIColors
}

type fieldData struct {
	FieldName string
	TypeName  string
}
