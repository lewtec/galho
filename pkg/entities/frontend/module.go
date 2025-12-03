package frontend

import (
	"fmt"
	"path/filepath"

	"galho/pkg/core"
)

type FrontendModule struct {
	path string
	name string
	// In a real implementation, we might need to find the related GraphQL schema
	// for Relay compilation. For now, I'll assume a convention or a flag.
	// Based on the spec: --schema ./internal/crm/api/schema.graphql
	// This implies we might need to look for other modules or configuration.
}

func NewFrontendModule(path string) *FrontendModule {
	name := filepath.Base(filepath.Dir(path))

	dir := filepath.Dir(path)
	if filepath.Base(path) == "frontend" {
		name = filepath.Base(dir)
	}

	return &FrontendModule{
		path: path,
		name: name,
	}
}

func (m *FrontendModule) Type() string {
	return "frontend"
}

func (m *FrontendModule) Path() string {
	return m.path
}

func (m *FrontendModule) GenerateTasks() ([]core.Task, error) {
	taskName := fmt.Sprintf("gen:%s:frontend:relay", m.name)

	// TODO: Smarter schema detection.
	// For now, I will use a placeholder or look for a sibling 'api' directory.
	// Assuming structure internal/NAME/frontend and internal/NAME/api
	// Using relative path from the frontend directory to the api directory
	runCmd := `bun run relay-compiler`

	return []core.Task{
		{
			Name:        taskName,
			Description: fmt.Sprintf("Generate Relay code for %s", m.name),
			Run:         runCmd,
			Dir:         m.path,
		},
	}, nil
}
