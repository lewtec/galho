package mise

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/lewtec/galho/pkg/core"
	// Import entities to trigger init() and register finders
	_ "github.com/lewtec/galho/pkg/entities/database"
	_ "github.com/lewtec/galho/pkg/entities/frontend"
	_ "github.com/lewtec/galho/pkg/entities/graphql"

	toml "github.com/pelletier/go-toml/v2"
	"github.com/spf13/cobra"
)

func init() {
	core.RegisterGenerateCommand("tasks", &cobra.Command{
		Use:   "tasks",
		Short: "Generate mise tasks from project modules",
		Long:  "Scan discovered modules and write .mise/galho.toml with generation tasks",
		RunE: func(cmd *cobra.Command, args []string) error {
			return GenerateTasksToml("")
		},
	})
}

type MiseTask struct {
	Description string   `toml:"description,omitempty"`
	Run         string   `toml:"run,omitempty"`
	Depends     []string `toml:"depends,omitempty"`
	Dir         string   `toml:"dir,omitempty"`
}

// MiseConfig is the root shape of .mise/galho.toml (mise [tasks.*] tables).
type MiseConfig struct {
	Tasks map[string]MiseTask `toml:"tasks"`
}

// GenerateTasksToml scans the project and writes .mise/galho.toml under projectRoot.
// If projectRoot is empty, the discovered project directory is used.
func GenerateTasksToml(projectRoot string) error {
	project, err := core.GetProject()
	if err != nil {
		return err
	}
	if projectRoot == "" {
		projectRoot = project.Dir()
	}

	tasks := make(map[string]MiseTask)

	err = project.FindModules(func(found core.ModuleFound) bool {
		moduleTasks, err := found.Module.GenerateTasks()
		if err != nil {
			fmt.Printf("Error generating tasks for module %s: %v\n", found.Module.Path(), err)
			return true
		}

		for _, task := range moduleTasks {
			taskDir := task.Dir
			if filepath.IsAbs(taskDir) {
				if relDir, err := filepath.Rel(project.Dir(), taskDir); err == nil {
					taskDir = relDir
				}
			}

			tasks[task.Name] = MiseTask{
				Description: task.Description,
				Run:         task.Run,
				Depends:     task.Depends,
				Dir:         taskDir,
			}
		}
		return true
	})
	if err != nil {
		return err
	}

	tasks["gen"] = MiseTask{
		Description: "Generate all code",
		Depends:     []string{"gen:*"},
	}

	miseDir := filepath.Join(projectRoot, ".mise")
	if err := os.MkdirAll(miseDir, 0755); err != nil {
		return err
	}

	outputPath := filepath.Join(miseDir, "galho.toml")
	f, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer f.Close()

	// Encode under the "tasks" table so mise sees [tasks."name"] entries.
	if err := toml.NewEncoder(f).Encode(MiseConfig{Tasks: tasks}); err != nil {
		return err
	}

	fmt.Printf("Mise tasks generated at %s\n", outputPath)
	return nil
}
