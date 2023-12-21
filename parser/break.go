package parser

import (
	"github.com/NuruProgramming/NuruIDLE/ast"
	"github.com/NuruProgramming/NuruIDLE/token"
)

func (p *Parser) parseBreak() *ast.Break {
	stmt := &ast.Break{Token: p.curToken}
	for p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}
