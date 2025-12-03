package graphql

import (
	"fmt"
	"os"
	"path/filepath"
)

const gqlgenTemplate = `schema:
  - schema.graphql

exec:
  filename: generated.go
  package: api

model:
  filename: models_gen.go
  package: api

resolver:
  layout: follow-schema
  dir: .
  package: api
`

const schemaTemplate = `type Query {
  todos: [Todo!]!
}

type Mutation {
  createTodo(text: String!): Todo!
}

type Todo {
  id: ID!
  text: String!
  done: Boolean!
  user: User!
}

type User {
  id: ID!
  name: String!
}
`

const resolverTemplate = `package api

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{}
`

func Generate(path string) error {
	// Ensure directory exists
	if err := os.MkdirAll(path, 0755); err != nil {
		return err
	}

	// Create gqlgen.yml
	if err := os.WriteFile(filepath.Join(path, "gqlgen.yml"), []byte(gqlgenTemplate), 0644); err != nil {
		return err
	}

	// Create schema.graphql
	if err := os.WriteFile(filepath.Join(path, "schema.graphql"), []byte(schemaTemplate), 0644); err != nil {
		return err
	}

	// Create resolver.go
	if err := os.WriteFile(filepath.Join(path, "resolver.go"), []byte(resolverTemplate), 0644); err != nil {
		return err
	}

	fmt.Printf("GraphQL module generated at %s\n", path)
	return nil
}
