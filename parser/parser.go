package parser

import (
	"kwago.dev/monkey/ast"
	"kwago.dev/monkey/lexer"
	"kwago.dev/monkey/token"
)

type Parser struct {
	lexer     *lexer.Lexer
	currToken token.Token
	peekToken token.Token
}

func (p Parser) Errors() []string {
	return []string{}
}

func NewParser(lexer *lexer.Lexer) *Parser {
	parser := Parser{
		lexer: lexer,
	}

	parser.nextToken()
	parser.nextToken()

	return &parser
}

func (p *Parser) nextToken() {
	p.currToken = p.peekToken
	p.peekToken = p.lexer.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
