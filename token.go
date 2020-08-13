package main

type TokenType string

const (
	EOF      = "EOF"
	LPARN    = "("
	RPARN    = ")"
	PLUS     = "+"
	MINUS    = "-"
	ASTERISK = "*"
	SLASH    = "/"
)

type Token struct {
	Type    TokenType
	Literal string
}
