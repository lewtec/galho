package core

import (
	galhocore "github.com/lewtec/galho/pkg/core"
	"github.com/lewtec/galho/pkg/utils/scaffold"

	"github.com/spf13/cobra"
)

func init() {
	cmd := &cobra.Command{
		Use:   "app [path]",
		Short: "Generate a new galho app scaffold",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			path := args[0]
			return scaffold.InstallFS(path, Template)
		},
	}

	galhocore.RegisterGenerateCommand("app", cmd)
}
