package scanner

import (
	"go-lox/tokens"
	"strconv"
)

type Scanner struct {
	Source  []rune
	Tokens  []*tokens.Token
	start   int
	current int
	line    int
}

func NewScanner(source []rune) *Scanner {
	return &Scanner{
		Source: source,
		Tokens: []*tokens.Token{},
		line:   1,
	}
}

func (s *Scanner) IsAtEnd() bool {
	return len(s.Source) == s.current
}

func (s *Scanner) ScanTokens() []*tokens.Token {
	for !s.IsAtEnd() {
		s.start = s.current
		s.ScanToken()
	}
	s.Tokens = append(s.Tokens, &tokens.Token{
		Type:    tokens.EOF,
		Lexeme:  []rune{},
		Line:    s.line,
		Literal: nil,
	})

	return s.Tokens
}

func (s *Scanner) ScanToken() {
	c := s.advance()
	switch c {
	// Single tokens
	case '(':
		s.addToken(tokens.LeftParen, nil)
	case ')':
		s.addToken(tokens.RightParen, nil)
	case '{':
		s.addToken(tokens.LeftBrace, nil)
	case '}':
		s.addToken(tokens.RightBrace, nil)
	case ',':
		s.addToken(tokens.Comma, nil)
	case '.':
		s.addToken(tokens.Dot, nil)
	case '-':
		s.addToken(tokens.Minus, nil)
	case '+':
		s.addToken(tokens.Plus, nil)
	case ';':
		s.addToken(tokens.Semicolon, nil)
	case '*':
		s.addToken(tokens.Star, nil)

	// Multi tokens
	case '!':
		if s.match('=') {
			s.addToken(tokens.BangEqual, nil)
		} else {
			s.addToken(tokens.Bang, nil)
		}
	case '=':
		if s.match('=') {
			s.addToken(tokens.EqualEqual, nil)
		} else {
			s.addToken(tokens.Equal, nil)
		}
	case '<':
		if s.match('=') {
			s.addToken(tokens.LessEqual, nil)
		} else {
			s.addToken(tokens.Less, nil)
		}
	case '>':
		if s.match('=') {
			s.addToken(tokens.GreaterEqual, nil)
		} else {
			s.addToken(tokens.Greater, nil)
		}

	case '/':
		// Double slash is comment til EOL
		if s.match('/') {
			for s.peek() != '\n' && !s.IsAtEnd() {
				s.advance()
			}
		} else {
			s.addToken(tokens.Slash, nil)
		}

	// String literals
	case '"':
		s.string()

	// Whitespace cases
	case ' ':
	case '\r':
	case '\t':
	case '\n':
		s.line++

	// Reserved words
	case 'o':
		if s.peek() == 'r' {
			s.addToken(tokens.Or, nil)
		}

	default:
		if s.isDigit(c) {
			s.number()
		} else if s.isAlpha(c) {
			s.identifier()
		} else {
			panic("Unexpected character: " + string(s.line))
		}
	}
}

func (s *Scanner) isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func (s *Scanner) isAlpha(r rune) bool {
	return (r >= 'a' && r <= 'z') ||
		(r >= 'A' && r <= 'Z') ||
		r == '_'
}

func (s *Scanner) isAlphaNumeric(r rune) bool {
	return s.isAlpha(r) || s.isDigit(r)
}

func (s *Scanner) match(expected rune) bool {
	if s.IsAtEnd() {
		return false
	}
	if s.Source[s.current] != expected {
		return false
	}
	s.current++
	return true
}

func (s *Scanner) identifier() {
	for s.isAlphaNumeric(s.peek()) {
		s.advance()
	}

	// See if the identifier is a reserved word.
	text := s.Source[s.start:s.current]
	token := keywords[string(text)]
	if token == "" {
		token = tokens.Identifier
	}
	s.addToken(token, nil)
}

func (s *Scanner) string() {
	for s.peek() != '"' && !s.IsAtEnd() {
		if s.peek() == '\n' {
			s.line++
		}
		s.advance()
	}
	// Unterminated
	if s.IsAtEnd() {
		panic("Unterminated string: line " + string(s.line))
	}

	// The closing "
	s.advance()

	// Trim surrounding quotes
	value := s.Source[s.start+1 : s.current-1]
	s.addToken(tokens.String, value)
}

func (s *Scanner) number() {
	for s.isDigit(s.peek()) {
		s.advance()
	}

	// Fractional parts
	if s.peek() == '.' && s.isDigit(s.peekNext()) {
		// consume the decimal
		s.advance()

		for s.isDigit(s.peek()) {
			s.advance()
		}
	}

	substr := s.Source[s.start:s.current]
	number, err := strconv.ParseFloat(string(substr), 64)
	if err != nil {
		panic("Error parsing number: line " + string(s.line))
	}
	s.addToken(tokens.Number, number)
}

func (s *Scanner) peek() rune {
	if s.IsAtEnd() {
		return rune(0)
	}

	return s.Source[s.current]
}

func (s *Scanner) peekNext() rune {
	if s.IsAtEnd() {
		return rune(0)
	}

	return s.Source[s.current+1]
}

func (s *Scanner) advance() rune {
	s.current++
	return s.Source[s.current-1]
}

func (s *Scanner) addToken(t tokens.TokenType, literal interface{}) {
	text := s.Source[s.start:s.current]
	s.Tokens = append(s.Tokens, &tokens.Token{
		Type:    t,
		Lexeme:  text,
		Literal: literal,
		Line:    s.line,
	})
}
