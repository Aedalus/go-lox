package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "go-lox",
	Short:   "Go-lox is a golang implementation of Bob Nystrom's Lox language",
	Version: "0.0.1",
	// Run: func(cmd *cobra.Command, args []string) {
	// 	// Do Stuff Here
	// },
}

func Execute() {
	rootCmd.AddCommand(parseCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
