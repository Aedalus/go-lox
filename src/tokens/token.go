package tokens

import "fmt"

type Token struct {
	Type    TokenType
	Lexeme  []rune
	Literal interface{}
	Line    int
}

func NewToken(t TokenType, lex []rune, literal interface{}, line int) *Token {
	return &Token{
		Type:    t,
		Lexeme:  lex,
		Literal: literal,
		Line:    line,
	}
}

func (t *Token) String() string {
	return fmt.Sprintf("%s %s", t.Type, string(t.Lexeme))
}
