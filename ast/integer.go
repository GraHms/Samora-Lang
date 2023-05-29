package ast

import (
	"latin/token"
)

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) String() string       { return il.Token.Literal }
func (il *IntegerLiteral) expressionNode()      {}
