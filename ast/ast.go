package ast

import "kwago.dev/monkey/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}

	return ""
}

type LetStatement struct {
	Name  *Identifier
	Value Expression
	Token token.Token
}

func (l LetStatement) statementNode() {}

func (l LetStatement) TokenLiteral() string {
	return l.Token.Literal
}

// Identifier is an expression since `let x = valueProducingIdentifier` is possible
type Identifier struct {
	Token token.Token
	Value string
}

func (i Identifier) expressionNode() {}

func (i Identifier) TokenLiteral() string {
	return i.Token.Literal
}
