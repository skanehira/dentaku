package eval

import (
	"github.com/skanehira/dentaku/ast"
	"github.com/skanehira/dentaku/object"
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		return evalExpressions(node.Expressions)
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	case *ast.FloatLiteral:
		return &object.Float{Value: node.Value}
	case *ast.PrefixExpression:
		right := Eval(node.Right)
		return evalPrefixExpression(right)
	case *ast.InfixExpression:
		left := Eval(node.Left)
		right := Eval(node.Right)
		return evalInfixExpression(left, node.Operator, right)
	}
	return nil
}

func evalInfixExpression(left object.Object, operator string, right object.Object) object.Object {
	l, ok := left.(*object.Integer)
	if !ok {
		return nil
	}
	r, ok := right.(*object.Integer)
	if !ok {
		return nil
	}

	switch operator {
	case "*":
		return &object.Integer{Value: l.Value * r.Value}
	case "+":
		return &object.Integer{Value: l.Value + r.Value}
	case "/":
		return &object.Integer{Value: l.Value / r.Value}
	case "-":
		return &object.Integer{Value: l.Value - r.Value}
	}

	return nil
}

func evalPrefixExpression(obj object.Object) object.Object {
	switch v := obj.(type) {
	case *object.Integer:
		return &object.Integer{Value: -v.Value}
	case *object.Float:
		return &object.Float{Value: -v.Value}
	}
	return nil
}

func evalExpressions(exps []ast.Expression) object.Object {
	var obj object.Object

	for _, exp := range exps {
		obj = Eval(exp)
	}

	return obj
}
