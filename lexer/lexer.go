package lexer

import "kwago.dev/monkey/token"
import "errors"

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
		return token.Token{
			Type: token.EOF,
		}
	}

	switch char {
	case '=':
		return token.Assign
	case '+':
		return token.Plus
	case '(':
		return token.LParen
	case ')':
		return token.RParen
	case '{':
		return token.LBrace
	case '}':
		return token.RBrace
	case ',':
		return token.Comma
	case ';':
		return token.Semicolon
	default:
		return token.Token{
			Type:    token.ILLEGAL,
			Literal: string(char),
		}
	}
}

func (l *Lexer) nextChar() (byte, error) { // returning byte here means that we'll only support ASCII for now
	if l.position > len(l.input)-1 {
		return 0, errors.New("EOF") // TODO: How will this differentiate with an actual `0` literal
	} else {
		l.position += 1
		return l.input[l.position], nil
	}
}
