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

var (
	//go:embed templates/explain.tmpl
	templateExplainEmbed embed.FS
	templateExplain      *template.Template
)

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

		err := templateExplain.Execute(os.Stdout, explainData{
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

func init() {
	templateExplainContent, err := templateExplainEmbed.ReadFile("templates/explain.tmpl")
	if err != nil {
		panic(err)
	}
	templateExplain, err = template.New("explain").Funcs(sprig.FuncMap()).Parse(string(templateExplainContent))
	if err != nil {
		panic(err)
	}

	rootCmd.AddCommand(explainCmd)
}
