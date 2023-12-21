package repl

import (
	"github.com/NuruProgramming/NuruIDLE/evaluator"
	"github.com/NuruProgramming/NuruIDLE/lexer"
	"github.com/NuruProgramming/NuruIDLE/object"
	"github.com/NuruProgramming/NuruIDLE/parser"
)

type Output struct {
	Evaluated string
	Errors    []string
}

func Read(contents string, env *object.Environment) Output {

	l := lexer.New(contents)
	p := parser.New(l)

	var o Output

	program := p.ParseProgram()

	if len(p.Errors()) != 0 {
		o.Errors = p.Errors()
		return o
	}

	evaluated := evaluator.Eval(program, env)
	if evaluated != nil {
		if evaluated.Type() == object.ERROR_OBJ {
			o.Errors = []string{evaluated.Inspect()}
			return o
		}
		if evaluated.Type() != object.NULL_OBJ {
			o.Evaluated = evaluated.Inspect()
			return o
		}
	}

	return o
}

func Start(in string, env *object.Environment) Output {

	l := lexer.New(in)
	p := parser.New(l)

	var o Output

	program := p.ParseProgram()

	if len(p.Errors()) != 0 {
		o.Errors = p.Errors()
		return o
	}

	evaluated := evaluator.Eval(program, env)
	if evaluated != nil {
		if evaluated.Type() == object.ERROR_OBJ {
			o.Errors = []string{evaluated.Inspect()}
			return o
		}
		if evaluated.Type() != object.NULL_OBJ {
			o.Evaluated = evaluated.Inspect()
			return o
		}
	}

	return o
}
