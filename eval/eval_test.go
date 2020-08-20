package eval

import (
	"reflect"
	"testing"

	"github.com/skanehira/dentaku/lexer"
	"github.com/skanehira/dentaku/object"
	"github.com/skanehira/dentaku/parser"
)

func TestIntegerExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"5", 5},
		{"10", 10},
		{"-5", -5},
		{"-10", -10},
		{"10 + 4", 14},
		{"10 + 4", 14},
	}

	for _, tt := range tests {
		v := testEval(tt.input)
		testIntegerObject(t, v, tt.expected)
	}
}

func testEval(in string) object.Object {
	l := lexer.New(in)
	p := parser.New(l)

	pp := p.ParseProgram()
	return Eval(pp)
}

func testIntegerObject(t *testing.T, obj object.Object, expected interface{}) {
	result, ok := obj.(*object.Integer)
	if !ok {
		t.Fatalf("object not integer. got=%T (%+v)", obj, obj)
	}

	if reflect.DeepEqual(result.Value, expected) {
		t.Fatalf("object has wrong value. got=%d, want=%d", result.Value, expected)
	}
}
