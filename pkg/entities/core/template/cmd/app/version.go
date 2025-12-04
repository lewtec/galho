package main

import (
	version "GALHO/app"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "show app version",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println(strings.TrimSpace(version.Version))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
