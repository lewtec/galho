package main

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "Sample galho app",
}

func main() {
	// Cobra already prints Execute errors; only set the process exit code.
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
