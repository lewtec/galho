package frontend

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

	if err := templates.CopyDir("frontend", path); err != nil {
		return err
	}

	fmt.Printf("Frontend module generated at %s\n", path)
	return nil
}
