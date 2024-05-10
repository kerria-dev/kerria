// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Kerria

package cmd

import (
	"github.com/spf13/cobra"
	"k8s.io/klog/v2"
)

var (
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Display the application version",
		Run: func(_ *cobra.Command, _ []string) {
			klog.Infof("commit %s%s%s, built on %s%s%s\n",
				colors.VersionControl,
				commit,
				colors.Reset,
				colors.Timestamp,
				date,
				colors.Reset)
		},
	}
)

func init() {
	rootCmd.AddCommand(versionCmd)
}
