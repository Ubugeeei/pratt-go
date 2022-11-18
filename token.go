package main

type Token struct {
	_type   TokenType
	literal string
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
