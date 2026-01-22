package database

import (
	"fmt"
	"path/filepath"

	"github.com/lewtec/galho/pkg/core"
)

type DatabaseModule struct {
	path string
	name string
}

func NewDatabaseModule(path string) *DatabaseModule {
	// The module name is the name of the parent directory.
	// e.g. internal/crm/db -> crm
	// internal/db -> internal
	name := filepath.Base(filepath.Dir(path))

	return &DatabaseModule{
		path: path,
		name: name,
	}
}

func (m *DatabaseModule) Type() string {
	return "database"
}

func (m *DatabaseModule) Path() string {
	return m.path
}

func (m *DatabaseModule) Name() string {
	return m.name
}

func (m *DatabaseModule) GenerateTasks() ([]core.Task, error) {
	taskName := fmt.Sprintf("gen:%s:db:sqlc", m.name)

	return []core.Task{
		{
			Name:        taskName,
			Description: fmt.Sprintf("Generate SQLC code for %s", m.name),
			Run:         "sqlc generate -f sqlc.yaml",
			Dir:         m.path,
		},
	}, nil
}
