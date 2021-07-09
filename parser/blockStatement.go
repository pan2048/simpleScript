package parser

import (
	"simple/ast"
	"simple/token"
)

func (p *Parser) parseBlockStatement() ast.BlockStatement {
	block := ast.BlockStatement{
		Token:      p.curToken,
		Statements: []ast.Statement{},
	}
	p.nextToken()
	for !p.curTokenIs(token.RBRACE) && !p.curTokenIs(token.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			block.Statements = append(block.Statements, stmt)
		}
		p.nextToken()
	}
	return block
}
