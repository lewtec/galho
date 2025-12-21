package main

import (
	"fmt"
	"os"

	"github.com/lewtec/galho/cmd/galho/entities"
	generate_cmd "github.com/lewtec/galho/cmd/galho/generate"
	modules_cmd "github.com/lewtec/galho/cmd/galho/modules"
	"github.com/lewtec/galho/pkg/core"
	_ "github.com/lewtec/galho/pkg/entities/database"
	_ "github.com/lewtec/galho/pkg/entities/frontend"
	_ "github.com/lewtec/galho/pkg/entities/graphql"

	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "galho",
	Short: "Galho is a modular framework for Golang",
}

func init() {

	Command.AddCommand(generate_cmd.Command)
	Command.AddCommand(modules_cmd.Command)

	entities.AddEntityCommands(Command)

	// Populate generate commands
	for _, cmd := range core.GetGenerateCommands() {
		generate_cmd.Command.AddCommand(cmd)
	}

}

func main() {
	if err := Command.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
