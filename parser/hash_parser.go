package parser

import (
	"github.com/grahms/samoralang/ast"
	"github.com/grahms/samoralang/token"
)

func (p *Parser) parseHashLiteral() ast.Expression {
	h := &ast.HashLiteral{Token: p.curToken}
	h.Pairs = make(map[ast.Expression]ast.Expression)
	for !p.peekTokenIs(token.RBRACE) {
		p.nextToken()
		key := p.parseExpression(LOWEST)
		if !p.expectPeek(token.COLON) {
			return nil
		}
		p.nextToken()
		value := p.parseExpression(LOWEST)
		h.Pairs[key] = value
		if !p.peekTokenIs(token.RBRACE) && !p.expectPeek(token.COMMA) {
			return nil
		}
	}
	if !p.expectPeek(token.RBRACE) {
		return nil
	}
	return h
}
