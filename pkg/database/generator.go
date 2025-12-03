package database

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

	if err := templates.CopyDir("database", path); err != nil {
		return err
	}

	fmt.Printf("Database module generated at %s\n", path)
	return nil
}
