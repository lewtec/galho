package database

import (
	"github.com/lewtec/galho/pkg/core"
	"github.com/lewtec/galho/pkg/utils/scaffold"

	"github.com/spf13/cobra"
)

func init() {
	cmd := &cobra.Command{
		Use:   "database [path]",
		Short: "Generate a database module",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			path := args[0]
			return scaffold.InstallFS(path, Template)
		},
	}

	core.RegisterGenerateCommand("database", cmd)
}
