package parser

import (
	"errors"
	"fmt"
	"go-lox/ast"
	"go-lox/tokens"
)

type Parser struct {
	tokens  []*tokens.Token
	current int
}

func (p *Parser) Parse() (ast.Expr, error) {
	return p.expression()
}

func (p *Parser) expression() (ast.Expr, error) {
	return p.equality()
}

func (p *Parser) equality() (ast.Expr, error) {
	expr, err := p.comparison()
	if err != nil {
		return nil, err
	}

	for p.match(tokens.BangEqual, tokens.EqualEqual) {
		operator := p.previous()
		right, err := p.comparison()
		if err != nil {
			return nil, err
		}
		expr = &ast.BinaryExpr{
			Left:     expr,
			Operator: operator,
			Right:    right,
		}
	}
	return expr, nil
}

func (p *Parser) comparison() (ast.Expr, error) {
	expr, err := p.addition()
	if err != nil {
		return nil, err
	}

	for p.match(tokens.Greater, tokens.GreaterEqual, tokens.Less, tokens.LessEqual) {
		operator := p.previous()
		right, err := p.addition()
		if err != nil {
			return nil, err
		}
		expr = &ast.BinaryExpr{
			Left:     expr,
			Operator: operator,
			Right:    right,
		}
	}

	return expr, nil
}

func (p *Parser) addition() (ast.Expr, error) {
	expr, err := p.multiplication()
	if err != nil {
		return nil, err
	}
	for p.match(tokens.Minus, tokens.Plus) {
		operator := p.previous()
		right, err := p.multiplication()
		if err != nil {
			return nil, err
		}
		expr = &ast.BinaryExpr{
			Left:     expr,
			Operator: operator,
			Right:    right,
		}
	}
	return expr, nil
}

func (p *Parser) multiplication() (ast.Expr, error) {
	expr, err := p.unary()
	if err != nil {
		return nil, err
	}

	for p.match(tokens.Slash, tokens.Star) {
		operator := p.previous()
		right, err := p.unary()
		if err != nil {
			return nil, err
		}
		expr = &ast.BinaryExpr{
			Left:     expr,
			Operator: operator,
			Right:    right,
		}
	}
	return expr, nil
}

func (p *Parser) unary() (ast.Expr, error) {
	if p.match(tokens.Bang, tokens.Minus) {
		operator := p.previous()
		right, err := p.unary()
		if err != nil {
			return nil, err
		}
		return &ast.UnaryExpr{
			Operator: operator,
			Right:    right,
		}, nil
	}
	return p.primary()
}

func (p *Parser) primary() (ast.Expr, error) {
	if p.match(tokens.False) {
		return &ast.LiteralExpr{Value: false}, nil
	} else if p.match(tokens.True) {
		return &ast.LiteralExpr{Value: true}, nil
	} else if p.match(tokens.Nil) {
		return &ast.LiteralExpr{Value: nil}, nil
	} else if p.match(tokens.Number, tokens.String) {
		return &ast.LiteralExpr{Value: p.previous().Literal}, nil
	} else if p.match(tokens.LeftParen) {
		expr := p.expression()
		p.consume(tokens.RightParen, "Expect ')' after expression.")
		return &ast.GroupingExpr{Expression: expr}, nil
	} else {
		errMsg := fmt.Sprintf("%s Expect expression", p.peek())
		return nil, errors.New(errMsg)
	}
}

func (p *Parser) consume(t tokens.TokenType, message string) (*tokens.Token, error) {
	if p.check(t) {
		return p.advance(), nil
	}
	return nil, errors.New(message)
}

func (p *Parser) error(t *tokens.Token, message string) error {
	errStr := fmt.Sprintf("%s %+v", message, t)
	println(errStr)
	return errors.New(errStr)
}

func (p *Parser) synchronize() {
	p.advance()

	for !p.isAtEnd() {
		if p.previous().Type == tokens.Semicolon {
			return
		}

		switch p.peek().Type {
		case tokens.Class:
			fallthrough
		case tokens.Fun:
			fallthrough
		case tokens.Var:
			fallthrough
		case tokens.For:
			fallthrough
		case tokens.If:
			fallthrough
		case tokens.While:
			fallthrough
		case tokens.Print:
			fallthrough
		case tokens.Return:
			return
		}

		p.advance()
	}
}

func (p *Parser) match(types ...tokens.TokenType) bool {
	for _, t := range types {
		if p.check(t) {
			p.advance()
			return true
		}
	}
	return false
}

func (p *Parser) check(t tokens.TokenType) bool {
	if p.isAtEnd() {
		return false
	}
	return p.peek().Type == t
}

func (p *Parser) advance() *tokens.Token {
	if !p.isAtEnd() {
		p.current++
	}
	return p.previous()
}

func (p *Parser) isAtEnd() bool {
	return p.peek().Type == tokens.EOF
}

func (p *Parser) peek() *tokens.Token {
	return p.tokens[p.current]
}

func (p *Parser) previous() *tokens.Token {
	return p.tokens[p.current-1]
}
