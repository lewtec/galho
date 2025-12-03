package graphql

import (
	"fmt"
	"path/filepath"

	"galho/pkg/core"
)

type GraphQLModule struct {
	path string
	name string
}

func NewGraphQLModule(path string) *GraphQLModule {
	name := filepath.Base(filepath.Dir(path))
	// Try to get the parent module name
	// e.g. internal/crm/api -> crm

	dir := filepath.Dir(path)
	if filepath.Base(path) == "api" {
		name = filepath.Base(dir)
	}

	return &GraphQLModule{
		path: path,
		name: name,
	}
}

func (m *GraphQLModule) Type() string {
	return "graphql"
}

func (m *GraphQLModule) Path() string {
	return m.path
}

func (m *GraphQLModule) GenerateTasks() ([]core.Task, error) {
	taskName := fmt.Sprintf("gen:%s:api:gqlgen", m.name)
	configPath := filepath.Join(m.path, "gqlgen.yml")

	return []core.Task{
		{
			Name:        taskName,
			Description: fmt.Sprintf("Generate GraphQL code for %s", m.name),
			Run:         fmt.Sprintf("gqlgen generate -c %s", configPath),
			Dir:         m.path,
		},
	}, nil
}
