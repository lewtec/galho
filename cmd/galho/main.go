package main

import (
	"fmt"
	"os"

	"galho/pkg/core"
	entities "galho/pkg/entities/core"
	_ "galho/pkg/entities/database"
	_ "galho/pkg/entities/frontend"
	_ "galho/pkg/entities/graphql"
	"galho/pkg/utils/scaffold"

	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "galho",
	Short: "Galho is a modular framework for Golang",
}

func init() {
	// 1. Setup 'init' command
	initCmd := &cobra.Command{
		Use:   "init [path]",
		Short: "Initialize a new Galho project",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			path := "."
			if len(args) > 0 {
				path = args[0]
			}
			return scaffold.InstallFS(path, entities.Template)
		},
	}
	Command.AddCommand(initCmd)

	// 2. Setup 'generate' command
	generateCmd := &cobra.Command{
		Use:   "generate",
		Short: "Generate components for the project",
	}

	// Add registered generate subcommands
	for _, cmd := range core.GetGenerateCommands() {
		generateCmd.AddCommand(cmd)
	}
	Command.AddCommand(generateCmd)

	// 3. Setup entity commands (e.g. db, etc.)
	for _, entityCmd := range core.GetEntityCommands() {
		Command.AddCommand(entityCmd.Command)
	}
}

func main() {
	if err := Command.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
