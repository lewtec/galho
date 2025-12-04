package main

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "Sample galho app",
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatalf("error: %s", err)
		os.Exit(1)
	}
}
