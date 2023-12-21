package parser

import (
	"fmt"
	"strconv"

	"github.com/NuruProgramming/NuruIDLE/ast"
)

func (p *Parser) parseFloatLiteral() ast.Expression {
	fl := &ast.FloatLiteral{Token: p.curToken}
	value, err := strconv.ParseFloat(p.curToken.Literal, 64)
	if err != nil {
		msg := fmt.Sprintf("Hatuwezi kuparse %q kama desimali", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}
	fl.Value = value
	return fl
}
