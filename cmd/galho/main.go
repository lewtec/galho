package main

import (
	"fmt"
	"galho/pkg/entities/core"
	"galho/pkg/entities/database"
	"galho/pkg/entities/frontend"
	"galho/pkg/entities/graphql"
	"galho/pkg/utils/mise"
	"galho/pkg/utils/scaffold"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "galho",
		Short: "Galho is a modular framework for Golang",
	}

	var initCmd = &cobra.Command{
		Use:   "init",
		Short: "Initialize a new project in the current directory",
		RunE: func(cmd *cobra.Command, args []string) error {
			workingDir, err := os.Getwd()
			if err != nil {
				return fmt.Errorf("error getting work directory: %w", err)
			}
			scaffold.InstallFS(workingDir, core.Template)

			fmt.Println("Project initialized successfully.")
		},
	}

	var generateCmd = &cobra.Command{
		Use:   "generate",
		Short: "Generate code or configuration",
	}

	var generateDatabaseCmd = &cobra.Command{
		Use:   "database [path]",
		Short: "Generate a database module",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			path := args[0]

			return scaffold.InstallFS(path, database.Template)
		},
	}

	var generateGraphQLCmd = &cobra.Command{
		Use:   "graphql [path]",
		Short: "Generate a graphql module",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			path := args[0]

			return scaffold.InstallFS(path, graphql.Template)
		},
	}

	var generateFrontendCmd = &cobra.Command{
		Use:   "frontend [path]",
		Short: "Generate a frontend module",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			path := args[0]
			return scaffold.InstallFS(path, frontend.Template)
		},
	}

	var generateTasksCmd = &cobra.Command{
		Use:   "tasks",
		Short: "Generate mise tasks",
		Run: func(cmd *cobra.Command, args []string) {
			cwd, err := os.Getwd()
			if err != nil {
				fmt.Printf("Error getting current working directory: %v\n", err)
				os.Exit(1)
			}
			if err := mise.GenerateTasksToml(cwd); err != nil {
				fmt.Printf("Error generating tasks: %v\n", err)
				os.Exit(1)
			}
		},
	}

	generateCmd.AddCommand(generateDatabaseCmd)
	generateCmd.AddCommand(generateGraphQLCmd)
	generateCmd.AddCommand(generateFrontendCmd)
	generateCmd.AddCommand(generateTasksCmd)

	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(generateCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
