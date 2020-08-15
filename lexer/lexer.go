package lexer

import (
	"github.com/skanehira/dentaku/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
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
	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[pos:l.position]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\r' || l.ch == '\n' {
		l.readChar()
	}
}

func (l *Lexer) peekCharIs(ch byte) bool {
	if len(l.input) <= l.readPosition {
		return false
	}

	if l.input[l.readPosition] == ch {
		l.readChar()
		return true
	}
	return false
}

func (l *Lexer) peekChar() byte {
	if len(l.input) <= l.readPosition {
		return 0
	}

	return l.input[l.readPosition]
}

func (l *Lexer) preChar() byte {
	if len(l.input) <= l.position || l.position < 1 {
		return 0
	}
	return l.input[l.position-1]
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		// -1 + 2
		// 2 + -1
		if !isDigit(l.preChar()) && isDigit(l.peekChar()) {
			l.readChar()
			tok = l.newNumberToken()
			tok.Literal = "-" + tok.Literal
			return tok
		} else {
			tok = newToken(token.MINUS, l.ch)
		}
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '(':
		tok = newToken(token.LPARN, l.ch)
	case ')':
		tok = newToken(token.RPARN, l.ch)
	case 0:
		tok.Type = token.EOF
		return tok
	default:
		if isDigit(l.ch) {
			return l.newNumberToken()
		}
		tok.Type = token.ILLEGAL
		return tok
	}

	l.readChar()
	return tok
}

func (l *Lexer) newNumberToken() token.Token {
	var tok token.Token
	tok.Literal = l.readNumber()
	if l.ch == '.' {
		l.readChar()
		tok.Literal += "." + l.readNumber()
		tok.Type = token.FLOAT
	} else {
		tok.Type = token.INT
	}
	return tok
}

func newToken(tokenType token.TokenType, literal byte) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(literal),
	}
}
