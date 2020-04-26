package ast

import "go-lox/tokens"

type Expr interface{}

type AssignExpr struct {
	Name  *tokens.Token
	Value Expr
}

type BinaryExpr struct {
	Left     Expr
	Operator *tokens.Token
	Right    Expr
}

type CallExpr struct {
	Callee    Expr
	Paren     *tokens.Token
	Arguments []Expr
}

type GetExpr struct {
	Object Expr
	Name   *tokens.Token
}

type GroupingExpr struct {
	Expression Expr
}

type LiteralExpr struct {
	Value interface{}
}

type LogicalExpr struct {
	Left     Expr
	Operator *tokens.Token
	Right    Expr
}

type SetExpr struct {
	Object Expr
	Name   *tokens.Token
	Value  Expr
}

type SuperExpr struct {
	Keyword *tokens.Token
	Method  *tokens.Token
}

type ThisExpr struct {
	Keyword *tokens.Token
}

type UnaryExpr struct {
	Operator *tokens.Token
	Right    Expr
}

type VariableExpr struct {
	Name *tokens.Token
}
