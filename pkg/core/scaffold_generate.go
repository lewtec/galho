package core

import (
	"github.com/spf13/cobra"
)

// NewScaffoldGenerateCommand builds a `generate` subcommand that installs a
// template at the single path argument (ExactArgs(1)).
//
// install is typically a closure over scaffold.InstallFS(path, Template).
func NewScaffoldGenerateCommand(use, short string, install func(path string) error) *cobra.Command {
	return &cobra.Command{
		Use:   use,
		Short: short,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return install(args[0])
		},
	}
}
