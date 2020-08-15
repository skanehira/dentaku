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

	ILLEGAL = "ILLEGAL"
)

type Token struct {
	Type    TokenType
	Literal string
}
