package frontend

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/lewtec/galho/pkg/core"
)

func init() {
	core.RegisterModuleFinder("frontend", FindFrontendModules)
}

// FindFrontendModules searches for frontend modules in the project.
// A frontend module is identified by the presence of both package.json and App.tsx files.
func FindFrontendModules(p *core.Project) ([]core.Module, error) {
	var modules []core.Module

	err := filepath.Walk(p.Dir(), func(path string, info os.FileInfo, err error) error {
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

		// Check for package.json and verify App.tsx exists in the same directory
		if info.Name() == "package.json" {
			dir := filepath.Dir(path)
			if _, err := os.Stat(filepath.Join(dir, "App.tsx")); err == nil {
				modules = append(modules, NewFrontendModule(dir))
			}
		}

		return nil
	})

	return modules, err
}
