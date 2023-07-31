package ast

import (
	"bytes"
	"github.com/grahms/samoralang/token"
)

// ForLoopExpression holds a for-loop
type ForLoopExpression struct {
	// Token is the actual token
	Token token.Token

	// Condition is the expression used to determine if the loop
	// is still running.
	Condition Expression

	// Consequence is the set of statements to be executed for the
	// loop body.
	Consequence *BlockStatement
}

func (fle *ForLoopExpression) expressionNode() {}

// TokenLiteral returns the literal token.
func (fle *ForLoopExpression) TokenLiteral() string { return fle.Token.Literal }

// String returns this object as a string.
func (fle *ForLoopExpression) String() string {
	var out bytes.Buffer
	out.WriteString("for (")
	out.WriteString(fle.Condition.String())
	out.WriteString(" ) {")
	out.WriteString(fle.Consequence.String())
	out.WriteString("}")
	return out.String()
}
