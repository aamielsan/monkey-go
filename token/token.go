package token

// source -> lexer(source) -> tokens -> parser(tokens) -> AST

type TokenType string

const (
	ILLEGAL TokenType = "ILLEGAL"
	EOF     TokenType = "EOF"

	// Identifiers + literals
	IDENT TokenType = "IDENT" // identifier
	INT   TokenType = "INT"   // ints

	// Operators
	ASSIGN TokenType = "="
	PLUS   TokenType = "+"

	// Delimiters
	COMMA     TokenType = ","
	SEMICOLON TokenType = ";"

	LPAREN TokenType = "("
	RPAREN TokenType = ")"
	LBRACE TokenType = "{"
	RBRACE TokenType = "}"

	// Keywords
	FUNCTION TokenType = "FUNCTION"
	LET      TokenType = "LET"
)

type Token struct {
	Type    TokenType
	Literal string
}

func New(tokenType TokenType, literal string) Token {
	return Token{
		Type:    tokenType,
		Literal: literal,
	}
}

// Constant types
var EndOfFile = New(EOF, string(EOF))
var Assign = New(ASSIGN, string(ASSIGN))
var Plus = New(PLUS, string(PLUS))
var Comma = New(COMMA, string(COMMA))
var Semicolon = New(SEMICOLON, string(SEMICOLON))
var LParen = New(LPAREN, string(LPAREN))
var RParen = New(RPAREN, string(RPAREN))
var LBrace = New(LBRACE, string(LBRACE))
var RBrace = New(RBRACE, string(RBRACE))
