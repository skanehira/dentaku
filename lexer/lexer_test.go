package lexer

import (
	"testing"

	"github.com/skanehira/dentaku/token"
)

func TestNextToken(t *testing.T) {
	input := `5 + 10 - 2 * ( 4.4 / 2.2 )`

	l := New(input)

	tests := []struct {
		tokenType token.TokenType
		literal   string
	}{
		{token.INT, "5"},
		{token.PLUS, "+"},
		{token.INT, "10"},
		{token.MINUS, "-"},
		{token.INT, "2"},
		{token.ASTERISK, "*"},
		{token.LPARN, "("},
		{token.FLOAT, "4.4"},
		{token.SLASH, "/"},
		{token.FLOAT, "2.2"},
		{token.RPARN, ")"},
		{token.EOF, ""},
	}

	for i, tt := range tests {
		tok := l.NextToken()
		if tt.tokenType != tok.Type {
			t.Fatalf("tests[%d] unexpected token type: want=%q, got=%q", i, tt.tokenType, tok.Type)
		}

		if tt.literal != tok.Literal {
			t.Fatalf("tests[%d] unexpected token literal: want=%q, got=%q", i, tt.literal, tok.Literal)
		}
	}
}
