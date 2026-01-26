package graphql

import (
	"fmt"

	"github.com/lewtec/galho/pkg/core"
)

type GraphQLModule struct {
	core.BaseModule
}

func NewGraphQLModule(path string) *GraphQLModule {
	name := core.DeriveModuleName(path)

	return &GraphQLModule{
		BaseModule: core.NewBaseModule(path, name),
	}
}

func (m *GraphQLModule) Type() string {
	return "graphql"
}

func (m *GraphQLModule) GenerateTasks() ([]core.Task, error) {
	taskName := fmt.Sprintf("gen:%s:api:gqlgen", m.Name())

	return []core.Task{
		{
			Name:        taskName,
			Description: fmt.Sprintf("Generate GraphQL code for %s", m.Name()),
			Run:         "gqlgen generate -c gqlgen.yml",
			Dir:         m.Path(),
		},
	}, nil
}
