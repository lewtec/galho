package database

import (
	"os"
	"path/filepath"

	"github.com/lewtec/galho/pkg/core"
)

func init() {
	core.RegisterModuleFinder("database", FindDatabaseModules)
}

// FindDatabaseModules searches for database modules in the project.
// A database module is identified by the presence of a sqlc.yaml file.
func FindDatabaseModules(p *core.Project) ([]core.Module, error) {
	return core.WalkModules(p, func(path string, info os.FileInfo) (core.Module, error) {
		if info.Name() == "sqlc.yaml" {
			return NewDatabaseModule(filepath.Dir(path)), nil
		}
		return nil, nil
	})
}
