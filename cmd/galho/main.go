package main

import (
	"fmt"
	"os"

	"github.com/lewtec/galho/pkg/core"
	_ "github.com/lewtec/galho/pkg/entities/database"
	_ "github.com/lewtec/galho/pkg/entities/frontend"
	_ "github.com/lewtec/galho/pkg/entities/graphql"
	// Registers `generate tasks` (mise.GenerateTasksToml).
	_ "github.com/lewtec/galho/pkg/utils/mise"

	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "galho",
	Short: "Galho is a modular framework for Golang",
}

func init() {
	// Wire generate subcommands registered by entity packages (blank imports above).
	generateCmd := &cobra.Command{
		Use:   "generate",
		Short: "Generate code for entities",
	}
	for _, cmd := range core.GetGenerateCommands() {
		generateCmd.AddCommand(cmd)
	}
	Command.AddCommand(generateCmd)

	// List detected modules in the current project.
	modulesCmd := &cobra.Command{
		Use:   "modules",
		Short: "List modules in the project",
		RunE: func(cmd *cobra.Command, args []string) error {
			project, err := core.GetProject()
			if err != nil {
				return err
			}
			return project.FindModules(func(m core.ModuleFound) bool {
				fmt.Printf("[%s] %s (%s)\n", m.Finder, m.Module.Name(), m.Module.Path())
				return true
			})
		},
	}
	Command.AddCommand(modulesCmd)

	// Entity commands (db, etc.) registered via core.RegisterEntityCommand.
	for _, config := range core.GetEntityCommands() {
		Command.AddCommand(config.Command)
	}
}

func main() {
	if err := Command.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
