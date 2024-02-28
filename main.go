package main

import (
	"cgen/generator"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {

	var rootCmd = &cobra.Command{
		Use:   "cgen",
		Short: "Generate codebase for various languages",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Usage: cgen [language]")
		},
	}

	generator.Init(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
