package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"go-lox/scanner"
)

var parseCmd = &cobra.Command{
	Use:   "parse [file]",
	Short: "Parses a file and prints the AST",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]
		absPath, err := filepath.Abs(path)
		if err != nil {
			println("Could not resolve path")
			os.Exit(1)
		}

		file, err := ioutil.ReadFile(absPath)
		if err != nil {
			println("Error: Could not open file " + absPath)
			os.Exit(1)
		}

		s := scanner.NewScanner([]rune(string(file)))
		tokens := s.ScanTokens()

		for _, t := range tokens {
			fmt.Println(t.String())
		}
	},
}
