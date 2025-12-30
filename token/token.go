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

	//연산자
	ASSIGN = "="
	PLUS = "+"
	MINUS = "-"
	BANG = "!"
	ASTERISK = "*"
	SLASH = "/"
	LT = "<"
	GT = ">"
	COMMA = ","
	SEMICOLON = ";"
	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"
	EQ = "=="
	NOT_EQ = "!="

	//예약어
	FUNCTION = "FUNCTION"
	LET = "LET"
	TRUE = "TRUE"
	FALSE = "FALSE"
	IF = "IF"
	ELSE = "ELSE"
	RETURN = "RETURN"
)

// 문자열을 키, TokentType을 값으로 가지는 토큰테이블
// fn, let가 key로 들어오면 예약어에 맞는 
var keywords = map[string]TokenType {
	"fn": FUNCTION,
	"let": LET,
	"if": IF,
	"else": ELSE,
	"true": TRUE,
	"false": FALSE,
	"return": RETURN,
}

// 식별자를 읽었을때 이 식별자가 예약어인지 아닌지 검사
// 맞다면 예약어에 맞는 TokenType 상수를 반환
func LookupIdent(ident string) TokenType {
	// keywords맵에서 ident(string)를 찾아라, 있으면 ok가 true, FUNCTION 또는 LET 반환
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	// 예약어가 아니면 모두 식별자이므로 타입 IDENT 반환
	return IDENT
}