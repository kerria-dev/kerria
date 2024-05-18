// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Kerria

package cmd

import (
	"github.com/kerria-dev/kerria/cmd/util"
	"github.com/kerria-dev/kerria/pkg/openapi"
	"github.com/mattn/go-isatty"
	"github.com/spf13/cobra"
	"k8s.io/klog/v2"
	"os"
	kyoapi "sigs.k8s.io/kustomize/kyaml/openapi"
)

var (
	// version, commit, and date are set via build flags
	version = "dev"
	commit  = "none"
	date    = "unknown"
	// enableColors is set with a cobra flag and TTY detection
	enableColors bool
	// colors is initialized with a persistent pre-run based on the status of enableColors
	colors util.ANSIColors
)

var rootCmd = &cobra.Command{
	Use:   "kerria",
	Short: "✿ Kerria is an in-tree manifest management tool",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		klog.InitFlags(nil)

		colors = util.Colors(enableColors)
		klog.Infof("%s✿ Kerria %s%s\n", colors.KerriaFlower, version, colors.Reset)

		kyoapi.SuppressBuiltInSchemaUse()
		openapi.ApplyGlobalSchema()
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		klog.Flush()
	},
}

func Execute() {
	if !isatty.IsTerminal(os.Stdout.Fd()) && !isatty.IsCygwinTerminal(os.Stdout.Fd()) {
		enableColors = false
	}

	err := rootCmd.Execute()
	if err != nil {
		klog.Error(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&enableColors, "colors", true, "Enable color output")
}
