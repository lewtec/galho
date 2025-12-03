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

func (m *GraphQLModule) Name() string {
	return m.name
}

func (m *GraphQLModule) GenerateTasks() ([]core.Task, error) {
	taskName := fmt.Sprintf("gen:%s:api:gqlgen", m.name)

	return []core.Task{
		{
			Name:        taskName,
			Description: fmt.Sprintf("Generate GraphQL code for %s", m.name),
			Run:         "gqlgen generate -c gqlgen.yml",
			Dir:         m.path,
		},
	}, nil
}
