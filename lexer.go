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
	case '1':
		return Token{
			_type:   Number,
			literal: "1",
		}
	case '2':
		return Token{
			_type:   Number,
			literal: "2",
		}

	case ' ':
		return Token{
			_type:   Number,
			literal: "2",
		}

	default:
		return Token{
			_type:   ILLEGAL,
			literal: string(lx.ch),
		}
	}
}
