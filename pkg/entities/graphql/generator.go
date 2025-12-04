package graphql

import (
	"galho/pkg/core"
	"galho/pkg/utils/scaffold"

	"github.com/spf13/cobra"
)

func init() {
	cmd := &cobra.Command{
		Use:   "graphql [path]",
		Short: "Generate a graphql module",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			path := args[0]
			return scaffold.InstallFS(path, Template)
		},
	}

	core.RegisterGenerateCommand("graphql", cmd)
}
