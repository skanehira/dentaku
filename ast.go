package main

import (
	"bytes"
	"strconv"
)

type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
}

type Expression interface {
	Node
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

func (p *Program) String() string {
	var out bytes.Buffer
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

type ExpressionStatement struct {
	Token      Token
	Expression Expression
}

func (e *ExpressionStatement) TokenLiteral() string {
	return e.Token.Literal
}

func (e *ExpressionStatement) String() string {
	if e.Expression != nil {
		return e.Expression.String()
	}
	return ""
}

type Int struct {
	Token Token
	Value int64
}

func (i Int) String() string {
	return strconv.FormatInt(i.Value, 10)
}

func (i Int) TokenLiteral() string {
	return i.Token.Literal
}

type Float struct {
	Token Token
	Value float64
}

func (f Float) String() string {
	return strconv.FormatFloat(f.Value, 'f', -1, 64)
}

func (f Float) TokenLiteral() string {
	return f.Token.Literal
}
