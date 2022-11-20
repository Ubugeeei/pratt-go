package main

type Token struct {
	type_   TokenType
	literal string
}

type TokenType int

const (
	IllegalToken TokenType = iota
	EOF
	NumberToken
	PlusToken
	MinusToken
	SlashToken
	AsteriskToken
)

func (t Token) GetPrecedence() int {
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
