package main

import (
	"fmt"
	"os"

	"galho/cmd/galho/entities"
	generate_cmd "galho/cmd/galho/generate"
	modules_cmd "galho/cmd/galho/modules"
	_ "galho/pkg/entities/database"
	_ "galho/pkg/entities/frontend"
	_ "galho/pkg/entities/graphql"

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

}

func main() {
	if err := Command.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
