package token

type TokenType string
type Token struct{
	Type TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"

	//식별자
	IDENT = "IDENT"
	INT = "INT"

	ASSIGN = "="
	PLUS = "+"
	COMMA = ","
	SEMICOLON = ";"
	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	//예약어
	FUNCTION = "FUNCTION"
	LET = "LET"
)