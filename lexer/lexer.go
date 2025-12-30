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

// 반환값이 token.Token, 즉 구조체이고, 매개변수는 *Lexer 타입을 가짐
// 다음 토큰을 참조하는 함수
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

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
	case '!':
		tok = newToken(token.BANG, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isletter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}
	
	l.readChar() // 다음 토큰을 읽어들임
	return tok
}

// 매개변수로 들어오는 토큰을 받아서 처리
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// 식별자를 읽는 함수
// NextToken()함수는 문자를 1개씩 읽어들임. 하지만 식별자는 이름이기 때문에 이어서 읽어야 함
func (l *Lexer) readIdentifier() string {
	position := l.position// 식별자가 시작되는 위치를 저장

	// isletter함수로 알파벳과 '-'가 아닐때까지 문자열을 읽는다.
	for isletter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// 식별자의 들어갈 수 있는 문자를 선별하는 함수
// 따라서 Next_char() 과 같은 이름이 성립가능하게 한다.
// 만약 abc!!와 같은 이름을 허용하고 싶으면 조건문에 슬쩍 끼워넣으면 됨.
func isletter(ch byte) bool {
	return  'a' <= ch && ch <= 'z' ||
			'A' <= ch && ch <= 'Z' ||
	  		ch == '_'
}

// 공백 스킵 함수
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// 숫자를 읽는 함수
func (l *Lexer) readNumber() string {
	position := l.position

	// 입력 문자열에서 숫자가 있으면 읽는다
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// 다음 입력을 살펴보기만 하고 l.position, l.readPosition을 움직이지 않는다
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}