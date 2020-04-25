package tokens

type TokenType string

const (
	// Single Character Tokens
	LeftParen  TokenType = "LeftParen"
	RightParen TokenType = "RightParen"
	LeftBrace  TokenType = "LeftBrace"
	RightBrace TokenType = "RightBrace"
	Comma      TokenType = "Comma"
	Dot        TokenType = "Dot"
	Minus      TokenType = "Minus"
	Plus       TokenType = "Plus"
	Semicolon  TokenType = "Semicolon"
	Slash      TokenType = "Slash"
	Star       TokenType = "Star"

	// One or two character tokens
	Bang         TokenType = "Bang"
	BangEqual    TokenType = "BangEqual"
	Equal        TokenType = "Equal"
	EqualEqual   TokenType = "EqualEqual"
	Greater      TokenType = "Greater"
	GreaterEqual TokenType = "GreaterEqual"
	Less         TokenType = "Less"
	LessEqual    TokenType = "LessEqual"

	// Literals
	Identifier TokenType = "Identifier"
	String     TokenType = "String"
	Number     TokenType = "Number"

	// Keywords
	And    TokenType = "And"
	Class  TokenType = "Class"
	Else   TokenType = "Else"
	False  TokenType = "False"
	Fun    TokenType = "Fun"
	For    TokenType = "For"
	If     TokenType = "If"
	Nil    TokenType = "Nil"
	Or     TokenType = "Or"
	Print  TokenType = "Print"
	Return TokenType = "Return"
	Super  TokenType = "Super"
	This   TokenType = "This"
	True   TokenType = "True"
	Var    TokenType = "Var"
	While  TokenType = "While"

	EOF TokenType = "EOF"
)
