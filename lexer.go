package main

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func NewLexer(input string) *Lexer {
	l := &Lexer{
		input: input,
	}

	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition++
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readNumber() string {
	pos := l.position
	for !isDigit(l.ch) {
		l.readChar()
	}

	return l.input[pos:l.position]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\r' || l.ch == '\n' {
		l.readChar()
	}
}

func (l *Lexer) NextToken() Token {
	var tok Token

	l.skipWhitespace()

	switch l.ch {
	case '+':
		newToken(PLUS, l.ch)
	case '-':
		newToken(MINUS, l.ch)
	case '/':
		newToken(SLASH, l.ch)
	case '*':
		newToken(ASTERISK, l.ch)
	case '(':
		newToken(LPARN, l.ch)
	case ')':
		newToken(RPARN, l.ch)
	case 0:
		tok.Type = EOF
		return tok
	default:
		tok.Type = ILLEGAL
		return tok
	}

	l.readChar()

	return tok
}

func newToken(tokenType TokenType, literal byte) Token {
	return Token{
		Type:    tokenType,
		Literal: string(literal),
	}
}
