package parser

import (
	"github.com/NuruProgramming/NuruIDLE/ast"
)

func (p *Parser) parseStringLiteral() ast.Expression {
	return &ast.StringLiteral{Token: p.curToken, Value: p.curToken.Literal}
}
