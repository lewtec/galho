package core

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

// CommandContext holds the resolved module and project for a command execution
type CommandContext struct {
	Project *Project
	Module  Module
}

// Context keys
type contextKey string

const (
	commandContextKey contextKey = "command_context"
)

// SetCommandContext stores the command context in the cobra command's context.
//
// cmd.Context() must already be set. cobra.Command.Execute / ExecuteContext
// seed the root context before PreRun and Run, so entity PersistentPreRun
// hooks rely on that instead of inventing context.Background here.
// Tests that call this without Execute must cmd.SetContext first (e.g. t.Context()).
func SetCommandContext(cmd *cobra.Command, ctx *CommandContext) {
	cmd.SetContext(context.WithValue(cmd.Context(), commandContextKey, ctx))
}

// GetCommandContext retrieves the command context from the cobra command
func GetCommandContext(cmd *cobra.Command) (*CommandContext, error) {
	if cmd.Context() == nil {
		return nil, fmt.Errorf("command context not initialized")
	}

	ctx := cmd.Context().Value(commandContextKey)
	if ctx == nil {
		return nil, fmt.Errorf("command context not found")
	}

	cmdCtx, ok := ctx.(*CommandContext)
	if !ok {
		return nil, fmt.Errorf("invalid command context type")
	}

	return cmdCtx, nil
}
