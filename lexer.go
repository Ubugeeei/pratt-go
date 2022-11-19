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
	for lx.peekChar() == ' ' || lx.peekChar() == '\t' || lx.peekChar() == '\n' || lx.peekChar() == '\r' {
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
	return '0' <= ch && ch <= '9' || ch == '.'
}

func (lx *Lexer) nextToken() Token {
	lx.skipWhitespace()
	lx.readChar()
	switch lx.ch {
	case '+':
		return Token{
			type_:   Plus,
			literal: "+",
		}
	case '-':
		return Token{
			type_:   Minus,
			literal: "-",
		}
	case '*':
		return Token{
			type_:   Asterisk,
			literal: "*",
		}
	case '/':
		return Token{
			type_:   Slash,
			literal: "/",
		}
	case '%':
		return Token{
			type_:   Percent,
			literal: "%",
		}
	case '1', '2', '3', '4', '5', '6', '7', '8', '9', '0':
		literal := lx.readDigit()
		return Token{
			type_:   Number,
			literal: literal,
		}
	case 0:
		return Token{
			type_:   EOF,
			literal: "",
		}
	default:
		return Token{
			type_:   ILLEGAL,
			literal: string(lx.ch),
		}
	}
}
