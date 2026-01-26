package frontend

import (
	"fmt"

	"github.com/lewtec/galho/pkg/core"
)

type FrontendModule struct {
	core.BaseModule
}

func NewFrontendModule(path string) *FrontendModule {
	name := core.DeriveModuleName(path)

	return &FrontendModule{
		BaseModule: core.NewBaseModule(path, name),
	}
}

func (m *FrontendModule) Type() string {
	return "frontend"
}

func (m *FrontendModule) GenerateTasks() ([]core.Task, error) {
	taskName := fmt.Sprintf("gen:%s:frontend:relay", m.Name())

	// TODO: Smarter schema detection.
	// For now, I will use a placeholder or look for a sibling 'api' directory.
	// Assuming structure internal/NAME/frontend and internal/NAME/api
	// Using relative path from the frontend directory to the api directory
	runCmd := `bun run relay-compiler`

	return []core.Task{
		{
			Name:        taskName,
			Description: fmt.Sprintf("Generate Relay code for %s", m.Name()),
			Run:         runCmd,
			Dir:         m.Path(),
		},
	}, nil
}
