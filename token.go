package main

type Token struct {
	type_   TokenType
	literal string
}

func (t Token) getPrecedence() int {
	switch t.type_ {
	case Plus:
		return 1
	case Minus:
		return 1
	case Asterisk:
		return 2
	case Slash:
		return 2
	case Percent:
		return 2
	default:
		return 0
	}
}

type TokenType int

const (
	ILLEGAL TokenType = iota
	EOF
	Number

	Plus
	Minus
	Slash
	Asterisk
	Percent
)
