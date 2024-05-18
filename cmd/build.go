// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Kerria

package cmd

import (
	"github.com/kerria-dev/kerria/pkg/build"
	"github.com/kerria-dev/kerria/pkg/processor"
	"github.com/kerria-dev/kerria/pkg/resources"
	"github.com/spf13/cobra"
	"k8s.io/klog/v2"
)

var (
	buildCmd = &cobra.Command{
		Use:   "build",
		Short: "Build the current Kerria-managed Repository",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Load configuration and lock
			repo, err := resources.ReadRepository()
			if err != nil {
				return err
			}
			lock, err := resources.ReadLockfile()
			if err != nil {
				return err
			}

			// Find intersection
			intersection := build.Intersect(repo, lock)

			// Pre-build processors
			message := processor.NewRepositoryMessage(repo, intersection)
			for _, proc := range repo.Processors {
				if proc.Stage == resources.StagePreBuild {
					klog.Infof("Executing pre-build processor %s", proc.Name)
					message.WithProcessor(proc)
					err = processor.DockerCommand(proc, message)
					if err != nil {
						return err
					}
				}
			}

			// Reconcile/Build
			klog.Info("Reconciling...")
			rebuilt, err := build.ReconcileDifferences(repo, lock, intersection)

			// Post-build processors
			message.WithRebuilt(rebuilt)
			for _, proc := range repo.Processors {
				if proc.Stage == resources.StagePostBuild {
					klog.Infof("Executing post-build processor %s", proc.Name)
					message.WithProcessor(proc)
					err = processor.DockerCommand(proc, message)
					if err != nil {
						return err
					}
				}
			}
			if err != nil {
				return err
			}

			// Write lock
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
