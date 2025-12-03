package core

// Module represents a detectable module in the project.
type Module interface {
	// Type returns the type of the module (e.g., "database", "graphql", "frontend").
	Type() string
	// Path returns the path to the module relative to the project root.
	Path() string
	// GenerateTasks returns the mise tasks for this module.
	GenerateTasks() ([]Task, error)
}

// Task represents a mise task.
type Task struct {
	Name        string
	Description string
	Run         string
	Dir         string
	Depends     []string
}

// Generator defines the interface for creating new modules.
type Generator interface {
	Generate(path string, name string) error
}
