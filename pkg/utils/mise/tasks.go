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
)

type MiseTask struct {
	Description string   `toml:"description,omitempty"`
	Run         string   `toml:"run,omitempty"`
	Depends     []string `toml:"depends,omitempty"`
	Dir         string   `toml:"dir,omitempty"`
}

type MiseConfig struct {
	Tasks map[string]MiseTask `toml:"tasks"`
}

// GenerateTasksToml scans the project and generates the .mise/tasks.toml file.
func GenerateTasksToml(projectRoot string) error {
	project, err := core.GetProject()
	if err != nil {
		return err
	}

	config := make(map[string]MiseTask)

	err = project.FindModules(func(found core.ModuleFound) bool {
		tasks, err := found.Module.GenerateTasks()
		if err != nil {
			// Log error but continue processing other modules
			fmt.Printf("Error generating tasks for module %s: %v\n", found.Module.Path(), err)
			return true
		}

		for _, task := range tasks {
			// Convert task.Dir to relative path if it's absolute
			taskDir := task.Dir
			if filepath.IsAbs(taskDir) {
				relDir, err := filepath.Rel(project.Dir(), taskDir)
				if err == nil {
					taskDir = relDir
				}
			}

			config[task.Name] = MiseTask{
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

	// Add gen:all task
	config["gen"] = MiseTask{
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

	encoder := toml.NewEncoder(f)
	// Mise often prefers multiline strings for 'run' if they are long, but standard toml encoder
	// might handle it its own way. go-toml v2 is compliant.
	if err := encoder.Encode(config); err != nil {
		return err
	}

	fmt.Printf("Mise tasks generated at %s\n", outputPath)
	return nil
}
