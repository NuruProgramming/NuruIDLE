package evaluator

import "github.com/NuruProgramming/NuruIDLE/object"

func evalMinusPrefixOperatorExpression(right object.Object, line int) object.Object {
	switch obj := right.(type) {

	case *object.Integer:
		return &object.Integer{Value: -obj.Value}

	case *object.Float:
		return &object.Float{Value: -obj.Value}

	default:
		return newError("Operesheni Haieleweki: -%s", right.Type())
	}
}
func evalPlusPrefixOperatorExpression(right object.Object, line int) object.Object {
	switch obj := right.(type) {

	case *object.Integer:
		return &object.Integer{Value: obj.Value}

	case *object.Float:
		return &object.Float{Value: obj.Value}

	default:
		return newError("Operesheni Haieleweki: +%s", right.Type())
	}
}

func evalPrefixExpression(operator string, right object.Object, line int) object.Object {
	switch operator {
	case "!":
		return evalBangOperatorExpression(right)
	case "-":
		return evalMinusPrefixOperatorExpression(right, line)
	case "+":
		return evalPlusPrefixOperatorExpression(right, line)
	default:
		return newError("Operesheni Haieleweki: %s%s", operator, right.Type())
	}
}
