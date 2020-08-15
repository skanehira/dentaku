package token

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

	INT   = "INT"
	FLOAT = "FLOAT"
)

type Token struct {
	Type    TokenType
	Literal string
}
