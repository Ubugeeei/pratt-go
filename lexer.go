package main

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func NewLexer(input string) *Lexer {
	l := new(Lexer)
	l.input = input
	l.position = -1
	l.readPosition = 0
	l.ch = 0
	return l
}

func (lx *Lexer) readChar() {
	if lx.readPosition >= len(lx.input) {
		lx.ch = 0
	} else {
		lx.ch = lx.input[lx.readPosition]
	}
	lx.position = lx.readPosition
	lx.readPosition += 1
}

func (lx *Lexer) peekChar() byte {
	if lx.readPosition >= len(lx.input) {
		return 0
	}
	return lx.input[lx.readPosition]
}

func (lx *Lexer) backChar() {
	lx.position -= 1
	lx.readPosition -= 1
	lx.ch = lx.input[lx.position]
}

func (lx *Lexer) skipWhitespace() {
	for lx.ch == ' ' || lx.ch == '\t' || lx.ch == '\n' || lx.ch == '\r' {
		lx.readChar()
	}
}

func (l *Lexer) readDigit() string {
	position := l.position
	for l.isDigit(l.ch) {
		l.readChar()
	}
	res := l.input[position:l.position]
	l.backChar()
	return res
}

func (_ Lexer) isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (lx *Lexer) nextToken() Token {
	lx.readChar()
	switch lx.ch {
	case '+':
		return Token{
			_type:   Plus,
			literal: "+",
		}
	case '-':
		return Token{
			_type:   Minus,
			literal: "-",
		}
	case '*':
		return Token{
			_type:   Asterisk,
			literal: "*",
		}
	case '/':
		return Token{
			_type:   Slash,
			literal: "/",
		}
	case '%':
		return Token{
			_type:   Percent,
			literal: "%",
		}
		// TODO:
	case '1', '2', '3', '4', '5', '6', '7', '8', '9', '0':
		literal := lx.readDigit()
		return Token{
			_type:   Number,
			literal: literal,
		}

	default:
		return Token{
			_type:   ILLEGAL,
			literal: string(lx.ch),
		}
	}
}
