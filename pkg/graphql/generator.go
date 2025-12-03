package graphql

import (
	"fmt"
	"os"
	"galho/pkg/templates"
)

func Generate(path string) error {
	// Ensure directory exists
	if err := os.MkdirAll(path, 0755); err != nil {
		return err
	}

	if err := templates.CopyDir("graphql", path); err != nil {
		return err
	}

	fmt.Printf("GraphQL module generated at %s\n", path)
	return nil
}
