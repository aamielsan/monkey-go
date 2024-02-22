package token

// source -> lexer(source) -> tokens -> parser(tokens) -> AST

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL TokenType = "ILLEGAL"
	EOF     TokenType = "EOF"

	// Identifiers + literals
	IDENT TokenType = "IDENT" // identifier
	INT   TokenType = "INT"   // ints

	// Operators
	ASSIGN TokenType = "ASSIGN"
	PLUS   TokenType = "PLUS"

	// Delimiters
	COMMA     TokenType = "COMMA"
	SEMICOLON TokenType = "SEMICOLON"

	LPAREN TokenType = "LPAREN"
	RPAREN TokenType = "RPAREN"
	LBRACE TokenType = "LBRACE"
	RBRACE TokenType = "RBRACE"

	// Keywords
	FUNCTION TokenType = "FUNCTION"
	LET      TokenType = "LET"
)

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func New(tokenType TokenType, literal string) Token {
	return Token{
		Type:    tokenType,
		Literal: literal,
	}
}

func LookupTypeFrom(str string) TokenType {
	if tokenType, ok := keywords[str]; ok {
		return tokenType
	}
	return IDENT
}

// Factory & constants
var (
	Eof       = New(EOF, "EOF")
	Assign    = New(ASSIGN, "=")
	Plus      = New(PLUS, "+")
	Comma     = New(COMMA, ",")
	Semicolon = New(SEMICOLON, ";")
	LParen    = New(LPAREN, "(")
	RParen    = New(RPAREN, ")")
	LBrace    = New(LBRACE, "{")
	RBrace    = New(RBRACE, "}")
	Let       = New(LET, "let")
)

func Illegal(literal string) Token {
	return New("ILLEGAL", literal)
}

func Ident(literal string) Token {
	return New("IDENT", literal)
}

func Int(literal string) Token {
	return New("INT", literal)
}

func Function(literal string) Token {
	return New("FUNCTIOn", literal)
}
