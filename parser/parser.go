package parser

import (
	"fmt"
	"strconv"

	"github.com/skanehira/dentaku/ast"
	"github.com/skanehira/dentaku/lexer"
	"github.com/skanehira/dentaku/token"
)

const (
	_ int = iota
	LOWSET
	SUM     // + or -
	PRODUCT // * or /
	PREFIX  // -10 or +10
)

type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
)

type Parser struct {
	l              *lexer.Lexer
	curToken       token.Token
	peekToken      token.Token
	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l: l,
	}

	p.prefixParseFns = make(map[token.TokenType]prefixParseFn)

	p.registerPrefix(token.INT, p.parseIntegerLiteral)
	p.registerPrefix(token.FLOAT, p.parseFloatLiteral)
	p.registerPrefix(token.MINUS, p.parsePrefixExpression)

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	for p.curToken.Type != token.EOF {
		exp := p.parseExpression(LOWSET)
		if exp != nil {
			program.Expressions = append(program.Expressions, exp)
		}
		p.nextToken()
	}
	return program
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixParseFns[p.curToken.Type]
	if prefix == nil {
		panic(fmt.Sprintf("no parse functions. token is %#+v", p.curToken))
	}
	leftExp := prefix()
	return leftExp
}

func (p *Parser) parsePrefixExpression() ast.Expression {
	expression := &ast.PrefixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
	}

	p.nextToken()

	expression.Right = p.parseExpression(PREFIX)
	return expression
}

func (p *Parser) parseIntegerLiteral() ast.Expression {
	lit := &ast.IntegerLiteral{
		Token: p.curToken,
	}

	v, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
	if err != nil {
		return nil
	}
	lit.Value = v

	return lit
}

func (p *Parser) parseFloatLiteral() ast.Expression {
	lit := &ast.FloatLiteral{
		Token: p.curToken,
	}

	v, err := strconv.ParseFloat(p.curToken.Literal, 64)
	if err != nil {
		return nil
	}
	lit.Value = v

	return lit
}

func (p *Parser) curTokenIs(token token.Token) bool {
	return p.curToken == token

}

func (p *Parser) peekTokenIs(token token.Token) bool {
	return p.peekToken == token
}

func (p *Parser) expectPeek(token token.Token) bool {
	if p.peekTokenIs(token) {
		p.nextToken()
		return true
	}
	return false
}

func (p *Parser) registerPrefix(token token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[token] = fn
}

func (p *Parser) registerInfix(token token.TokenType, fn infixParseFn) {
	p.infixParseFns[token] = fn
}
