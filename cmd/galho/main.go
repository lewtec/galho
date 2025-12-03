package main

import (
	"fmt"
	"os"
	"path/filepath"

	"galho/pkg/database"
	"galho/pkg/frontend"
	"galho/pkg/graphql"
	"galho/pkg/mise"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "galho",
		Short: "Galho is a modular framework for Golang",
	}

	var initCmd = &cobra.Command{
		Use:   "init [name]",
		Short: "Initialize a new project",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			projectName := args[0]
			fmt.Printf("Initializing project %s\n", projectName)

			// Create project directory
			if err := os.MkdirAll(projectName, 0755); err != nil {
				fmt.Printf("Error creating project directory: %v\n", err)
				os.Exit(1)
			}

			// Create .mise/tasks.toml
			if err := os.MkdirAll(filepath.Join(projectName, ".mise"), 0755); err != nil {
				fmt.Printf("Error creating .mise directory: %v\n", err)
				os.Exit(1)
			}

			// Create mise.toml
			miseToml := `_.file.includes = [".mise/tasks.toml"]

[tasks.dev]
run = "air"
`
			if err := os.WriteFile(filepath.Join(projectName, "mise.toml"), []byte(miseToml), 0644); err != nil {
				fmt.Printf("Error creating mise.toml: %v\n", err)
				os.Exit(1)
			}

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
		Run: func(cmd *cobra.Command, args []string) {
			path := args[0]
			if err := database.Generate(path); err != nil {
				fmt.Printf("Error generating database module: %v\n", err)
				os.Exit(1)
			}
		},
	}

	var generateGraphQLCmd = &cobra.Command{
		Use:   "graphql [path]",
		Short: "Generate a graphql module",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			path := args[0]
			if err := graphql.Generate(path); err != nil {
				fmt.Printf("Error generating graphql module: %v\n", err)
				os.Exit(1)
			}
		},
	}

	var generateFrontendCmd = &cobra.Command{
		Use:   "frontend [path]",
		Short: "Generate a frontend module",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			path := args[0]
			if err := frontend.Generate(path); err != nil {
				fmt.Printf("Error generating frontend module: %v\n", err)
				os.Exit(1)
			}
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
