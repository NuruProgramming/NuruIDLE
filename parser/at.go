package parser

import "github.com/NuruProgramming/NuruIDLE/ast"

func (p *Parser) parseAt() ast.Expression {
	return &ast.At{Token: p.curToken}
}
