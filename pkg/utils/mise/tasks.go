package mise

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"galho/pkg/core"
	"galho/pkg/entities/database"
	"galho/pkg/entities/frontend"
	"galho/pkg/entities/graphql"

	toml "github.com/pelletier/go-toml/v2"
)

type MiseTask struct {
	Description string   `toml:"description,omitempty"`
	Run         string   `toml:"run,omitempty"`
	Depends     []string `toml:"depends,omitempty"`
}

type MiseConfig struct {
	Tasks map[string]MiseTask `toml:"tasks"`
}

// GenerateTasksToml scans the project and generates the .mise/tasks.toml file.
func GenerateTasksToml(projectRoot string) error {
	modules, err := detectModules(projectRoot)
	if err != nil {
		return err
	}

	config := MiseConfig{
		Tasks: make(map[string]MiseTask),
	}

	for _, mod := range modules {
		tasks, err := mod.GenerateTasks()
		if err != nil {
			return err
		}

		for _, task := range tasks {
			config.Tasks[task.Name] = MiseTask{
				Description: task.Description,
				Run:         task.Run,
				Depends:     task.Depends,
			}
		}
	}

	// Add gen:all task
	config.Tasks["gen"] = MiseTask{
		Description: "Generate all code",
		Depends:     []string{"gen:*"},
	}

	miseDir := filepath.Join(projectRoot, ".mise")
	if err := os.MkdirAll(miseDir, 0755); err != nil {
		return err
	}

	outputPath := filepath.Join(miseDir, "tasks.toml")

	f, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer f.Close()

	encoder := toml.NewEncoder(f)
	// Mise often prefers multiline strings for 'run' if they are long, but standard toml encoder
	// might handle it its own way. go-toml v2 is compliant.
	if err := encoder.Encode(config); err != nil {
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
