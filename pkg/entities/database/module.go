package database

import (
	"fmt"
	"path/filepath"

	"github.com/lewtec/galho/pkg/core"
)

type DatabaseModule struct {
	core.BaseModule
}

func NewDatabaseModule(path string) *DatabaseModule {
	name := core.DeriveModuleName(path)

	// Fallback if path is just internal/db (parent is internal, but name derived as internal?)
	// Original logic: if name == "db" { name = "app" }
	// This applies when parent dir is "db".
	if name == "db" && filepath.Base(path) != "db" {
		name = "app"
	}

	return &DatabaseModule{
		BaseModule: core.NewBaseModule(path, name),
	}
}

func (m *DatabaseModule) Type() string {
	return "database"
}

func (m *DatabaseModule) GenerateTasks() ([]core.Task, error) {
	taskName := fmt.Sprintf("gen:%s:db:sqlc", m.Name())

	return []core.Task{
		{
			Name:        taskName,
			Description: fmt.Sprintf("Generate SQLC code for %s", m.Name()),
			Run:         "sqlc generate -f sqlc.yaml",
			Dir:         m.Path(),
		},
	}, nil
}
