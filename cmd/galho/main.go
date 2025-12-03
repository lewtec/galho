package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "galho",
	Short: "Galho is a modular framework for Golang",
}

func main() {

	if err := Command.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
