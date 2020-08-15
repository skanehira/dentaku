package main

const (
	_ int = iota
	LOWSET
	SUM     // + or -
	PRODUCT // * or /
	PREFIX  // -10 or +10
)

type Parser struct {
	l         *Lexer
	curToken  Token
	peekToken Token
}

func NewParser(l *Lexer) *Parser {
	p := &Parser{
		l: l,
	}

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.l.NextToken()
}

func (p *Parser) ParseProgram() *Program {
	program := &Program{}
	for p.curToken.Type != EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}
	return program
}

func (p *Parser) parseStatement() Statement {
	switch p.curToken {
	default:
		return nil
	}
}

func (p *Parser) curTokenIs(token Token) bool {
	return p.curToken == token

}

func (p *Parser) peekTokenIs(token Token) bool {
	return p.peekToken == token
}

func (p *Parser) expectPeek(token Token) bool {
	if p.peekTokenIs(token) {
		p.nextToken()
		return true
	}
	return false
}
