package main

import (
	"github.com/spf13/cobra"
	"log"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:     "godot",
		Short:   "Godot: toolkit for go micro",
		Long:    "Godot: toolkit for go micro",
		Version: "v0.0.1",
	}
	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
