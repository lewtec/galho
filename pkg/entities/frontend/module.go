package frontend

import (
	"fmt"

	"github.com/lewtec/galho/pkg/core"
)

type FrontendModule struct {
	path string
	name string
	// Schema discovery for Relay is still convention-based (sibling api/).
	// See GenerateTasks; smarter detection is intentionally deferred.
}

func NewFrontendModule(path string) *FrontendModule {
	return &FrontendModule{
		path: path,
		name: core.DeriveModuleName(path, "frontend"),
	}
}

func (m *FrontendModule) Type() string {
	return "frontend"
}

func (m *FrontendModule) Path() string {
	return m.path
}

func (m *FrontendModule) Name() string {
	return m.name
}

func (m *FrontendModule) GenerateTasks() ([]core.Task, error) {
	taskName := fmt.Sprintf("gen:%s:frontend:relay", m.name)

	// Relay schema is expected at ../api/schema.graphql (relay.config.json).
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
