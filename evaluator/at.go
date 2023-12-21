package evaluator

import (
	"github.com/NuruProgramming/NuruIDLE/ast"
	"github.com/NuruProgramming/NuruIDLE/object"
)

func evalAt(node *ast.At, env *object.Environment) object.Object {
	if at, ok := env.Get("@"); ok {
		return at
	}
	return newError("Iko nje ya scope")
}
