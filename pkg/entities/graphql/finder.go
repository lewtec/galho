package graphql

import (
	"os"
	"path/filepath"
	"strings"

	"galho/pkg/core"
)

func init() {
	core.RegisterModuleFinder("graphql", FindGraphQLModules)
}

// FindGraphQLModules searches for GraphQL modules in the project.
// A GraphQL module is identified by the presence of a gqlgen.yml file.
func FindGraphQLModules(p *core.Project) ([]core.Module, error) {
	var modules []core.Module

	err := filepath.Walk(p.Dir(), func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			// Skip .git, node_modules, etc
			if strings.HasPrefix(info.Name(), ".") || info.Name() == "node_modules" {
				return filepath.SkipDir
			}
			return nil
		}

		// Check for gqlgen.yml to identify GraphQL modules
		if info.Name() == "gqlgen.yml" {
			modules = append(modules, NewGraphQLModule(filepath.Dir(path)))
		}

		return nil
	})

	return modules, err
}
