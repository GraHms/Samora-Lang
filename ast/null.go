package ast

import "github.com/grahms/samoralang/token"

type NullLiteral struct {
	Token token.Token
}

func (n *NullLiteral) expressionNode() {}

func (n *NullLiteral) TokenLiteral() string { return n.Token.Literal }

func (n *NullLiteral) String() string { return n.Token.Literal }
