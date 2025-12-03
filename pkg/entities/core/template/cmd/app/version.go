package main

import "github.com/spf13/cobra"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "show app version",
	RunE: func(cmd *cobra.Command, args []string) error {
		println("version")

		//TODO: version logic
		return nil
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
