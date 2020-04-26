package ast

import (
	"fmt"
	"go-lox/tokens"
	"strings"
)

func parenthesize(name string, exprs ...Expr) string {
	var sb strings.Builder
	sb.WriteString("(")
	sb.WriteString(name)

	for _, e := range exprs {
		sb.WriteString(" ")
		sb.WriteString(Print(e))
	}
	sb.WriteString(")")

	return sb.String()
}

func parenthesize2(name string, parts ...interface{}) string {
	var sb strings.Builder
	sb.WriteString("(")
	sb.WriteString(name)
	for _, part := range parts {
		switch x := part.(type) {
		case *Expr:
			sb.WriteString(Print(x))
		case *Stmt:
			sb.WriteString(Print(x))
		case *tokens.Token:
			sb.WriteString(string(x.Lexeme))
		default:
			sb.WriteString(fmt.Sprintf("%v", x))
		}
	}
	sb.WriteString(")")
	return sb.String()
}

func Print(x interface{}) string {
	switch v := x.(type) {
	// Statements
	case *BlockStmt:
		var sb strings.Builder
		sb.WriteString("(block ")
		for _, s := range v.Statements {
			sb.WriteString(Print(s))
		}
		sb.WriteString(")")
		return sb.String()
	case *ClassStmt:
		var sb strings.Builder
		sb.WriteString("(class " + string(v.Name.Lexeme))
		if v.Superclass != nil {
			sb.WriteString(" < " + Print(v.Superclass))
		}
		for _, method := range v.Methods {
			sb.WriteString(" " + Print(method))
		}
		sb.WriteString(")")
		return sb.String()
	case *ExprStmt:
		return parenthesize(";", v.Expression)
	case *FunctionStmt:
		var sb strings.Builder
		sb.WriteString("(fun " + string(v.Name.Lexeme) + "(")
		for i, param := range v.Params {
			if i != 1 {
				sb.WriteString(" ")
			}
			sb.WriteString(string(param.Lexeme))
		}
		sb.WriteString(") ")
		for _, body := range v.Body {
			sb.WriteString(Print(body))
		}
		sb.WriteString(")")
		return sb.String()
	case *IfStmt:
		if v.ElseBranch == nil {
			return parenthesize2("if", v.Condition, v.ThenBranch)
		} else {
			return parenthesize2("if-else", v.Condition, v.ThenBranch, v.ElseBranch)
		}
	case *PrintStmt:
		return parenthesize("print", v.Expression)
	case *ReturnStmt:
		if v.Value == nil {
			return "(return)"
		}
		return parenthesize("return", v.Value)
	case *VariableStmt:
		if v.Initializer == nil {
			return parenthesize2("var", v.Name)
		}
		return parenthesize2("var", v.Name, "=", v.Initializer)
	case *WhileStmt:
		return parenthesize2("while", v.Condition, v.Body)

	// Expressions
	case *AssignExpr:
		return parenthesize2("=", v.Name.Lexeme, v.Value)
	case *BinaryExpr:
		return parenthesize(string(v.Operator.Lexeme), v.Left, v.Right)
	case *CallExpr:
		return parenthesize2("call", v.Callee, v.Arguments)
	case *GetExpr:
		return parenthesize2(".", v.Object, v.Name.Lexeme)
	case *GroupingExpr:
		return parenthesize("group", v.Expression)
	case *LiteralExpr:
		if v.Value == nil {
			return "nil"
		} else {
			return fmt.Sprintf("%v", v.Value)
		}
	case *LogicalExpr:
		return parenthesize(string(v.Operator.Lexeme), v.Left, v.Right)
	case *SetExpr:
		return parenthesize2("=", v.Object, v.Name.Lexeme, v.Value)
	case *SuperExpr:
		return parenthesize2("super", v.Method)
	case *ThisExpr:
		return "this"
	case *UnaryExpr:
		return parenthesize(string(v.Operator.Lexeme), v.Right)
	case *VariableExpr:
		return string(v.Name.Lexeme)
	default:
		panic(fmt.Sprintf("UNKNOWN_TYPE %+v", x))
	}
}
