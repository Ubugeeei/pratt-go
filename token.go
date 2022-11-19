package main

type Token struct {
	type_   TokenType
	literal string
}

type TokenType int

const (
	EOF TokenType = iota
	NumberToken
	PlusToken
	MinusToken
	SlashToken
	AsteriskToken
)

func (t Token) getPrecedence() int {
	switch t.type_ {
	case PlusToken:
		return 1
	case MinusToken:
		return 1
	case AsteriskToken:
		return 2
	case SlashToken:
		return 2
	default:
		return 0
	}
}
