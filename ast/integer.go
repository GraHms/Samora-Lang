package ast

import (
	"fmt"
	"github.com/grahms/samoralang/token"
)

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) String() string       { return il.Token.Literal }
func (il *IntegerLiteral) expressionNode()      {}

type FloatLiteral struct {
	Token token.Token // the '5.0', '5.00', etc. token
	Value float64
}

func (fl *FloatLiteral) TokenLiteral() string { return fl.Token.Literal }

// String returns a string representation of the float, formatted to 2 decimal places.
func (fl *FloatLiteral) String() string {
	return fmt.Sprintf("%.2f", fl.Value)
}
func (fl *FloatLiteral) expressionNode() {}
