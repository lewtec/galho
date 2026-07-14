package frontend

import (
	"os"
	"path/filepath"

	"github.com/lewtec/galho/pkg/core"
)

func init() {
	core.RegisterModuleFinder("frontend", FindFrontendModules)
}

// FindFrontendModules searches for frontend modules in the project.
// A frontend module is identified by the presence of both package.json and App.tsx files.
func FindFrontendModules(p *core.Project) ([]core.Module, error) {
	return core.WalkModules(p, func(path string, info os.FileInfo) (core.Module, error) {
		if info.Name() == "package.json" {
			dir := filepath.Dir(path)
			if _, err := os.Stat(filepath.Join(dir, "App.tsx")); err == nil {
				return NewFrontendModule(dir), nil
			}
		}
		return nil, nil
	})
}
