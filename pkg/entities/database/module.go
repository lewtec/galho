package database

import (
	"fmt"

	"github.com/lewtec/galho/pkg/core"
)

type DatabaseModule struct {
	path string
	name string
}

func NewDatabaseModule(path string) *DatabaseModule {
	return &DatabaseModule{
		path: path,
		name: core.DeriveModuleName(path, "db"),
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
