// Copyright 2021 Contributors to the Veraison project.
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var comidCmd = &cobra.Command{
	Use:   "comid",
	Short: "CoMID manipulation",

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help() // nolint: errcheck
			os.Exit(0)
		}
	},
}

func init() {
	rootCmd.AddCommand(comidCmd)
}
