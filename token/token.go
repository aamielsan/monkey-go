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
	ASSIGN   TokenType = "ASSIGN"
	PLUS     TokenType = "PLUS"
	MINUS    TokenType = "MINUS"
	SLASH    TokenType = "SLASH"
	BANG     TokenType = "BANG"
	ASTERISK TokenType = "ASTERISK"
	LT       TokenType = "LT"
	GT       TokenType = "GT"
	EQ       TokenType = "EQ"
	NOT_EQ   TokenType = "NOT_EQ"

	// Delimiters
	COMMA     TokenType = "COMMA"
	SEMICOLON TokenType = "SEMICOLON"

	// Boolean
	TRUE  TokenType = "TRUE"
	FALSE TokenType = "FALSE"

	LPAREN TokenType = "LPAREN"
	RPAREN TokenType = "RPAREN"
	LBRACE TokenType = "LBRACE"
	RBRACE TokenType = "RBRACE"

	// Keywords
	FUNCTION TokenType = "FUNCTION"
	LET      TokenType = "LET"
	IF       TokenType = "IF"
	ELSE     TokenType = "ELSE"
	RETURN   TokenType = "RETURN"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"if":     IF,
	"else":   ELSE,
	"==":     EQ,
	"!=":     NOT_EQ,
	"return": RETURN,
	"true":   TRUE,
	"false":  FALSE,
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
	Eof         = New(EOF, "")
	Assign      = New(ASSIGN, "=")
	Plus        = New(PLUS, "+")
	Minus       = New(MINUS, "-")
	Slash       = New(SLASH, "/")
	Bang        = New(BANG, "!")
	Asterisk    = New(ASTERISK, "*")
	Comma       = New(COMMA, ",")
	Semicolon   = New(SEMICOLON, ";")
	LParen      = New(LPAREN, "(")
	RParen      = New(RPAREN, ")")
	LBrace      = New(LBRACE, "{")
	RBrace      = New(RBRACE, "}")
	LessThan    = New(LT, "<")
	GreaterThan = New(GT, ">")
	Let         = New(LET, "let")
	True        = New(TRUE, "true")
	False       = New(FALSE, "false")
	Equal       = New(EQ, "==")
	NotEqual    = New(NOT_EQ, "!=")
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
