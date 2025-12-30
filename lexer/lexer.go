package lexer

import "monkey/token"

type Lexer struct{
	input string
	position int
	readPosition int
	ch byte
}

// 공용 함수(매개변수로 string타입을 받고 반환자료형은 *Lexer)
func New(input string) *Lexer {
	// Lexer 구조체를 하나 만들고 그 주소를 l에 저장
	// 인자로 받은 input에 값을 넣고, 나머지 구조체 필드는 0 초기화
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// Lexer 전용 메서드
func (l *Lexer) readChar() {

	// 다음 문자의 인덱스가 입력값의 길이보다 크면 안됨.
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		// 다음 문자의 값을 l.ch에 복사
		l.ch = l.input[l.readPosition]
	}

	// 다음 인덱스 참조
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token{
	var tok token.Token

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok  = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok  = newToken(token.LPAREN, l.ch)
	case ')':
		tok  = newToken(token.RPAREN, l.ch)
	case '{':
		tok  = newToken(token.LBRACE, l.ch)
	case '}':
		tok  = newToken(token.RBRACE, l.ch)			
	case ',':
		tok  = newToken(token.COMMA, l.ch)
	case '+':
		tok  = newToken(token.PLUS, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}

	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}