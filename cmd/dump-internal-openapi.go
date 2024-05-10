// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Kerria

package cmd

import (
	"encoding/json"
	"github.com/kerria-dev/kerria/pkg/openapi"
	"github.com/spf13/cobra"
	"os"
)

var dumpOpenAPIV2Cmd = &cobra.Command{
	Hidden: true,
	Use:    "dump-internal-openapiv2",
	Run: func(cmd *cobra.Command, args []string) {
		api := openapi.BuildInternalOpenAPIV2Specification()

		apiJSON, err := json.Marshal(api)
		if err != nil {
			panic(err)
		}

		_, err = os.Stdout.Write(apiJSON)
		if err != nil {
			panic(err)
		}
	},
}

var dumpOpenAPIV3Cmd = &cobra.Command{
	Hidden: true,
	Use:    "dump-internal-openapiv3",
	Run: func(cmd *cobra.Command, args []string) {
		api := openapi.BuildInternalOpenAPIV3Specification()

		apiJSON, err := json.Marshal(api)
		if err != nil {
			panic(err)
		}

		_, err = os.Stdout.Write(apiJSON)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(dumpOpenAPIV2Cmd)
	rootCmd.AddCommand(dumpOpenAPIV3Cmd)
}
