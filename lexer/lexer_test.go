package lexer

import (
	"kwago.dev/monkey/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	tests := []struct {
		exepectedType    token.TokenType
		exepectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
	}

	lexer := New(input)

	for idx, tt := range tests {
		token := lexer.NextToken()

		if token.Type != tt.exepectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, actual=%q", idx, tt.exepectedType, token.Type)
		}

		if token.Literal != tt.exepectedLiteral {
			t.Fatalf("test[%d] - literal wrong. expected=%q, actual=%q", idx, tt.exepectedLiteral, token.Literal)
		}
	}
}
