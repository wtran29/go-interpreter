package evaluator

import (
	"github.com/wtran29/go-interpreter/src/ast"
	"github.com/wtran29/go-interpreter/src/object"
)

// Eval takes in a ast.Node and returns an object.Object
func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	// Statements
	case *ast.Program:
		return evalStatements(node.Statements)

	case *ast.ExpressionStatement:
		return Eval(node.Expression)

	// Expressions
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	}

	return nil
}

// evalStatements evaluates each ast.Statement for an object.Object
func evalStatements(stmts []ast.Statement) object.Object {
	var result object.Object

	for _, statement := range stmts {
		result = Eval(statement)
	}

	return result
}
