package ast

import (
	"github.com/codytheroux96/go-interpreter/token"
)

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

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

type Identifier struct {
	Token token.Token
	Value string
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string { return ls.TokenLiteral()}

func (l *Identifier) expressionNode() {}
func (l *Identifier) TokenLiteral() string { return l.TokenLiteral()}