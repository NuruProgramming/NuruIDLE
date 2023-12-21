package evaluator

import (
	"github.com/NuruProgramming/NuruIDLE/ast"
	"github.com/NuruProgramming/NuruIDLE/object"
)

func evalIdentifier(node *ast.Identifier, env *object.Environment) object.Object {
	if val, ok := env.Get(node.Value); ok {
		return val
	}
	if builtin, ok := builtins[node.Value]; ok {
		return builtin
	}

	return newError("Neno Halifahamiki: %s", node.Value)
}
