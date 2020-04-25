package scanner

import (
	"go-lox/tokens"
)

var keywords = map[string]tokens.TokenType{
	"and":    tokens.And,
	"class":  tokens.Class,
	"else":   tokens.Else,
	"false":  tokens.False,
	"for":    tokens.For,
	"fun":    tokens.Fun,
	"if":     tokens.If,
	"nil":    tokens.Nil,
	"or":     tokens.Or,
	"print":  tokens.Print,
	"return": tokens.Return,
	"super":  tokens.Super,
	"this":   tokens.This,
	"true":   tokens.True,
	"var":    tokens.Var,
	"while":  tokens.While,
}
