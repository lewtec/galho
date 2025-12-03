package mise

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"galho/pkg/core"
	"galho/pkg/database"
	"galho/pkg/frontend"
	"galho/pkg/graphql"
)

// GenerateTasksToml scans the project and generates the .mise/tasks.toml file.
func GenerateTasksToml(projectRoot string) error {
	modules, err := detectModules(projectRoot)
	if err != nil {
		return err
	}

	var sb strings.Builder
	var allTaskNames []string

	for _, mod := range modules {
		tasks, err := mod.GenerateTasks()
		if err != nil {
			return err
		}

		for _, task := range tasks {
			allTaskNames = append(allTaskNames, fmt.Sprintf("%q", task.Name))

			sb.WriteString(fmt.Sprintf("[tasks.%q]\n", task.Name))
			if task.Description != "" {
				sb.WriteString(fmt.Sprintf("description = %q\n", task.Description))
			}

			// If run is multiline, use triple quotes
			if strings.Contains(task.Run, "\n") {
				sb.WriteString(fmt.Sprintf("run = \"\"\"\n%s\n\"\"\"\n", task.Run))
			} else {
				sb.WriteString(fmt.Sprintf("run = %q\n", task.Run))
			}

			if len(task.Depends) > 0 {
				sb.WriteString("depends = [")
				for i, dep := range task.Depends {
					if i > 0 {
						sb.WriteString(", ")
					}
					sb.WriteString(fmt.Sprintf("%q", dep))
				}
				sb.WriteString("]\n")
			}
			sb.WriteString("\n")
		}
	}

	// Add gen:all task
	sb.WriteString("[tasks.gen]\n")
	sb.WriteString("description = \"Generate all code\"\n")
	sb.WriteString("depends = [\"gen:*\"]\n") // Mise supports wildcards in depends, but explicit list is safer if not. Spec says depends = ["gen:*"]

	miseDir := filepath.Join(projectRoot, ".mise")
	if err := os.MkdirAll(miseDir, 0755); err != nil {
		return err
	}

	outputPath := filepath.Join(miseDir, "tasks.toml")
	if err := os.WriteFile(outputPath, []byte(sb.String()), 0644); err != nil {
		return err
	}

	fmt.Printf("Mise tasks generated at %s\n", outputPath)
	return nil
}

func detectModules(root string) ([]core.Module, error) {
	var modules []core.Module

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			// Skip .git, node_modules, etc
			if strings.HasPrefix(info.Name(), ".") || info.Name() == "node_modules" {
				return filepath.SkipDir
			}
			return nil
		}

		// Check for specific files to identify modules
		if info.Name() == "sqlc.yaml" {
			modules = append(modules, database.NewDatabaseModule(filepath.Dir(path)))
		} else if info.Name() == "gqlgen.yml" {
			modules = append(modules, graphql.NewGraphQLModule(filepath.Dir(path)))
		} else if info.Name() == "package.json" {
			// Check if it's a frontend module (e.g. has App.tsx)
			// This is a simplification. Spec says: "contém package.json + App.tsx"
			dir := filepath.Dir(path)
			if _, err := os.Stat(filepath.Join(dir, "App.tsx")); err == nil {
				modules = append(modules, frontend.NewFrontendModule(dir))
			}
		}

		return nil
	})

	return modules, err
}
