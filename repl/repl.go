package repl

import (
	"github.com/NuruProgramming/NuruIDLE/evaluator"
	"github.com/NuruProgramming/NuruIDLE/lexer"
	"github.com/NuruProgramming/NuruIDLE/object"
	"github.com/NuruProgramming/NuruIDLE/parser"
)

func Read(contents string, env *object.Environment) (string, []string) {

	l := lexer.New(contents)
	p := parser.New(l)

	program := p.ParseProgram()

	if len(p.Errors()) != 0 {
		return "", p.Errors()
	}

	evaluated := evaluator.Eval(program, env)
	if evaluated != nil {
		if evaluated.Type() != object.NULL_OBJ {
			return evaluated.Inspect(), nil
		}
	}

	return "", nil
}

func Start(in string, env *object.Environment) (string, []string) {

	l := lexer.New(in)
	p := parser.New(l)

	program := p.ParseProgram()

	if len(p.Errors()) != 0 {
		return "", p.Errors()
	}

	evaluated := evaluator.Eval(program, env)
	if evaluated != nil {
		if evaluated.Type() != object.NULL_OBJ {
			return evaluated.Inspect(), nil
		}
	}

	return "", nil
}
