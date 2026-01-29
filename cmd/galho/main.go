package main

import (
	"fmt"
	"os"

	"github.com/lewtec/galho/pkg/core"
	_ "github.com/lewtec/galho/pkg/entities/database"
	_ "github.com/lewtec/galho/pkg/entities/frontend"
	_ "github.com/lewtec/galho/pkg/entities/graphql"

	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "galho",
	Short: "Galho is a modular framework for Golang",
}

func init() {
	// Register Generate Command
	generateCmd := &cobra.Command{
		Use:   "generate",
		Short: "Generate code for entities",
	}
	for name, cmd := range core.GetGenerateCommands() {
		cmd.Use = name
		generateCmd.AddCommand(cmd)
	}
	Command.AddCommand(generateCmd)

	// Register Modules Command
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

	// Register Entity Commands
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
