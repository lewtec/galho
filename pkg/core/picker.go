package core

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// selectModuleInteractive presents an interactive picker for module selection
func selectModuleInteractive(modules []Module, entityType string) (Module, error) {
	// Use fzf if available, otherwise fall back to simple selection
	if hasFzf() {
		return selectWithFzf(modules, entityType)
	}
	return selectWithPrompt(modules, entityType)
}

// hasFzf checks if fzf is available
func hasFzf() bool {
	_, err := exec.LookPath("fzf")
	return err == nil
}

// selectWithFzf uses fzf for interactive selection
func selectWithFzf(modules []Module, entityType string) (Module, error) {
	// Create input for fzf
	var input strings.Builder
	for _, m := range modules {
		input.WriteString(fmt.Sprintf("%s\n", m.Path()))
	}

	// Run fzf
	cmd := exec.Command("fzf",
		"--prompt", fmt.Sprintf("Select %s module: ", entityType),
		"--height", "40%",
		"--reverse",
		"--border")

	cmd.Stdin = strings.NewReader(input.String())
	cmd.Stderr = os.Stderr

	output, err := cmd.Output()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok && exitErr.ExitCode() == 130 {
			return nil, fmt.Errorf("selection cancelled")
		}
		return nil, fmt.Errorf("fzf failed: %w", err)
	}

	selectedPath := strings.TrimSpace(string(output))

	// Find module by path
	for _, m := range modules {
		if m.Path() == selectedPath {
			return m, nil
		}
	}

	return nil, fmt.Errorf("invalid selection")
}

// selectWithPrompt uses simple numbered prompt for selection
func selectWithPrompt(modules []Module, entityType string) (Module, error) {
	fmt.Printf("Select %s module:\n", entityType)
	for i, m := range modules {
		fmt.Printf("  [%d] %s\n", i+1, m.Path())
	}
	fmt.Print("\nEnter number: ")

	var choice int
	_, err := fmt.Scanf("%d", &choice)
	if err != nil {
		return nil, fmt.Errorf("invalid input: %w", err)
	}

	if choice < 1 || choice > len(modules) {
		return nil, fmt.Errorf("invalid selection: %d", choice)
	}

	return modules[choice-1], nil
}
