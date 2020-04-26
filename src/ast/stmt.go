package ast

import "go-lox/tokens"

type Stmt interface {
}

type BlockStmt struct {
	Statements []*Stmt
}

type ClassStmt struct {
	Name       *tokens.Token
	Superclass *VariableExpr
	Methods    []*FunctionStmt
}

type ExprStmt struct {
	Expression *Expr
}

type FunctionStmt struct {
	Name   *tokens.Token
	Params []*tokens.Token
	Body   []*Stmt
}

type IfStmt struct {
	Condition  *Expr
	ThenBranch *Stmt
	ElseBranch *Stmt
}

type PrintStmt struct {
	Expression *Expr
}

type ReturnStmt struct {
	Keyword *tokens.Token
	Value   *Expr
}

type VariableStmt struct {
	Name        *tokens.Token
	Initializer *Expr
}

type WhileStmt struct {
	Condition *Expr
	Body      *Stmt
}
