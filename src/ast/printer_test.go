package ast

import (
	"go-lox/tokens"
	"testing"
)

func TestAST(t *testing.T) {
	expr := &BinaryExpr{
		Left: &UnaryExpr{
			Operator: &tokens.Token{
				Type:    tokens.Minus,
				Lexeme:  []rune{'-'},
				Literal: nil,
				Line:    1,
			},
			Right: &LiteralExpr{
				Value: 123,
			},
		},
		Operator: &tokens.Token{
			Type:    tokens.Star,
			Lexeme:  []rune{'*'},
			Literal: nil,
			Line:    1,
		},
		Right: &GroupingExpr{
			Expression: &LiteralExpr{
				Value: 45.67,
			},
		},
	}
	actual := Print(expr)
	expected := "(* (- 123) (group 45.67))"
	if actual != expected {
		t.Errorf(` got %s; expected %s;`, actual, expected)
	}

}
