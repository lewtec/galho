package database

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/lewtec/galho/pkg/core"
)

func init() {
	core.RegisterModuleFinder("database", FindDatabaseModules)
}

// FindDatabaseModules searches for database modules in the project.
// A database module is identified by the presence of a sqlc.yaml file.
func FindDatabaseModules(p *core.Project) ([]core.Module, error) {
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

		// Check for sqlc.yaml to identify database modules
		if info.Name() == "sqlc.yaml" {
			modules = append(modules, NewDatabaseModule(filepath.Dir(path)))
		}

		return nil
	})

	return modules, err
}
