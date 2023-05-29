package parser

import (
	"latin/ast"
	"strconv"
)

func (p *Parser) parseIntegerLiteral() ast.Expression {
	lit := &ast.IntegerLiteral{Token: p.curToken}
	value, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
	if err != nil {
		msg := "could not parse %q as integer"
		p.errors = append(p.errors, msg)
		return nil
	}
	lit.Value = value
	return lit
}
