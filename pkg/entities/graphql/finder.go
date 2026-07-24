package graphql

import (
	"os"
	"path/filepath"

	"github.com/lewtec/galho/pkg/core"
)

func init() {
	core.RegisterModuleFinder("graphql", FindGraphQLModules)
}

// FindGraphQLModules searches for GraphQL modules in the project.
// A GraphQL module is identified by the presence of a gqlgen.yml file.
func FindGraphQLModules(p *core.Project) ([]core.Module, error) {
	return core.WalkModules(p, func(path string, info os.FileInfo) (core.Module, error) {
		if info.Name() == "gqlgen.yml" {
			return NewGraphQLModule(filepath.Dir(path)), nil
		}
		return nil, nil
	})
}
