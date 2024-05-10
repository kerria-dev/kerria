// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Kerria

package cmd

import (
	"github.com/kerria-dev/kerria/pkg/build"
	"github.com/kerria-dev/kerria/pkg/resources"
	"github.com/spf13/cobra"
)

var (
	buildCmd = &cobra.Command{
		Use:   "build",
		Short: "Build the current Kerria-managed Repository",
		RunE: func(cmd *cobra.Command, args []string) error {
			repo, err := resources.ReadRepository()
			if err != nil {
				return err
			}
			lock, err := resources.ReadLockfile()
			if err != nil {
				return err
			}
			diff := build.CheckDifference(repo, lock)
			err = build.ReconcileDifferences(repo, lock, diff)
			if err != nil {
				return err
			}
			err = lock.Write()
			if err != nil {
				return err
			}
			return nil
		},
	}
)

func init() {
	rootCmd.AddCommand(buildCmd)
}
