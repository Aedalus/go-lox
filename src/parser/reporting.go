package parser

import (
	"fmt"
	"go-lox/tokens"
)

func report(line int, where string, message string) {
	fmt.Printf("[line %d] Error %s: %s", line, where, message)
}

func printError(t *tokens.Token, msg string) {
	if t.Type == tokens.EOF {
		report(t.Line, " at end", msg)
	} else {
		report(t.Line, " at '"+string(t.Lexeme)+"'", msg)
	}
}
