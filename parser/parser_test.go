package parser

import (
	"fmt"
	"testing"

	"github.com/skanehira/dentaku/ast"
	"github.com/skanehira/dentaku/lexer"
)

func TestParsingPrefixExpression(t *testing.T) {
	tests := []struct {
		input        string
		operator     string
		integerValue int64
	}{
		{"-1", "-", 1},
		{"-20", "-", 20},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)

		program := p.ParseProgram()

		if len(program.Expressions) != 1 {
			t.Fatalf("program.Expressions does not contain %d Expressions. got=%d\n", 1, len(program.Expressions))
		}

		exp, ok := program.Expressions[0].(*ast.PrefixExpression)
		if !ok {
			t.Fatalf("program.Expressions[0] is not ast.PrefixExpression. got=%T", program.Expressions[0])
		}

		if exp.Operator != tt.operator {
			t.Fatalf("exp.Operator is not %s. got=%s", tt.operator, exp.Operator)
		}

		if !testIntegerLiteral(t, exp.Right, tt.integerValue) {
			return
		}
	}
}

func testIntegerLiteral(t *testing.T, il ast.Expression, value int64) bool {
	t.Helper()
	integ, ok := il.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("il not *ast.IntegerLiteral. got=%T", il)
		return false
	}

	if integ.Value != value {
		t.Fatalf("integ.Value not %d. got=%d", value, integ.Value)
		return false
	}

	if integ.TokenLiteral() != fmt.Sprintf("%d", value) {
		t.Fatalf("integ.TokenLiteral() not %d. got=%s", value, integ.TokenLiteral())
		return false
	}
	return true
}
