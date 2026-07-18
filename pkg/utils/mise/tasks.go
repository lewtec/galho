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
//
// The output file is written atomically (temp file + rename) so a failed encode
// cannot leave a truncated live config. Module task generation errors fail the
// whole command instead of writing a partial task set.
func GenerateTasksToml(projectRoot string) error {
	project, err := core.GetProject()
	if err != nil {
		return err
	}
	if projectRoot == "" {
		projectRoot = project.Dir()
	}

	tasks := make(map[string]MiseTask)
	var taskErr error

	err = project.FindModules(func(found core.ModuleFound) bool {
		moduleTasks, err := found.Module.GenerateTasks()
		if err != nil {
			taskErr = fmt.Errorf("generating tasks for module %s: %w", found.Module.Path(), err)
			return false
		}

		for _, task := range moduleTasks {
			taskDir := task.Dir
			if filepath.IsAbs(taskDir) {
				// Rel against the output root so projectRoot overrides stay consistent.
				if relDir, err := filepath.Rel(projectRoot, taskDir); err == nil {
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
	if taskErr != nil {
		return taskErr
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
	if err := writeGalhoTomlAtomic(outputPath, MiseConfig{Tasks: tasks}); err != nil {
		return err
	}

	fmt.Printf("Mise tasks generated at %s\n", outputPath)
	return nil
}

// writeGalhoTomlAtomic encodes cfg to path via a same-dir temp file and rename
// so readers never see a half-written config.
func writeGalhoTomlAtomic(path string, cfg MiseConfig) error {
	dir := filepath.Dir(path)
	tmp, err := os.CreateTemp(dir, "galho-*.toml")
	if err != nil {
		return err
	}
	tmpName := tmp.Name()
	// Clean up the temp file if we fail before rename.
	defer os.Remove(tmpName)

	enc := toml.NewEncoder(tmp)
	if err := enc.Encode(cfg); err != nil {
		tmp.Close()
		return err
	}
	if err := tmp.Close(); err != nil {
		return err
	}

	return os.Rename(tmpName, path)
}
