package lexer

import (
	"errors"

	"kwago.dev/monkey/token"
)

// Lexer: lexer(source) => tokens
// Lexer takes in a source code as input and outputs a stream of tokens
type Lexer struct {
	input    string
	position int // current position
}

func New(source string) *Lexer {
	return &Lexer{
		input:    source,
		position: -1,
	}
}

func (l *Lexer) NextToken() token.Token {
	char, err := l.nextChar()
	if err != nil {
		return token.Eof
	}

	switch char {
	case '=':
		if peek, err := l.peekChar(); err == nil && peek == '=' {
			l.nextChar() // advance the cursor
			return token.Equal
		}
		return token.Assign
	case '+':
		return token.Plus
	case '-':
		return token.Minus
	case '*':
		return token.Asterisk
	case '/':
		return token.Slash
	case '(':
		return token.LParen
	case ')':
		return token.RParen
	case '{':
		return token.LBrace
	case '}':
		return token.RBrace
	case '<':
		return token.LessThan
	case '>':
		return token.GreaterThan
	case ',':
		return token.Comma
	case ';':
		return token.Semicolon
	case '!':
		if peek, err := l.peekChar(); err == nil && peek == '=' {
			l.nextChar() // advance the cursor
			return token.NotEqual
		}
		return token.Bang
	case ' ', '\t', '\n', '\r':
		return l.NextToken() // skip whitespaces
	default:
		if isLetter(char) {
			tokenLiteral := l.readIdentifier()
			tokenType := token.LookupTypeFrom(tokenLiteral)
			return token.New(tokenType, tokenLiteral)
		} else if isDigit(char) {
			tokenLiteral := l.readNumber()
			return token.Int(tokenLiteral)
		} else {
			return token.Illegal(string(char))
		}
	}
}

// returning byte here means that we'll only support ASCII for now
func (l *Lexer) nextChar() (byte, error) {
	l.position += 1
	if l.position > len(l.input)-1 {
		return 0, errors.New("EOF")
	} else {
		return l.input[l.position], nil
	}
}

func (l *Lexer) peekChar() (byte, error) {
	peekIdx := l.position + 1
	if peekIdx > len(l.input)-1 {
		return 0, errors.New("EOF")
	} else {
		return l.input[peekIdx], nil
	}
}

func (l *Lexer) readIdentifier() string {
	start := l.position

	for {
		peek, err := l.peekChar()
		if err != nil || !isLetter(peek) {
			break
		} else {
			l.nextChar()
		}
	}

	end := l.position + 1
	return l.input[start:end]
}

func (l *Lexer) readNumber() string {
	start := l.position

	for {
		peek, err := l.peekChar()
		if err != nil || !isDigit(peek) {
			break
		} else {
			l.nextChar()
		}
	}

	end := l.position + 1
	return l.input[start:end]
}

// use to test an potential identifier
// this means that our identifier can only start with letters or underscore
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(char byte) bool {
	return char >= '0' && char <= '9'
}
