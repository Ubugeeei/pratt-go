package main

import "errors"

type Parser struct {
	lx            *Lexer
	current_token Token
	peek_token    Token
}

func NewParser(lx *Lexer) *Parser {
	p := new(Parser)
	p.lx = lx
	first, _ := lx.GetNextToken()
	second, _ := lx.GetNextToken()
	p.current_token = first
	p.peek_token = second
	return p
}

func (p *Parser) Parse() (Node, error) {
	return p.parseInternal(0)
}

func (p *Parser) parseInternal(precedence int) (Node, error) {
	node, e := p.parsePrefix()
	if e != nil {
		return Node{}, e
	}

	for precedence < p.peek_token.GetPrecedence() && p.peek_token.type_ != EOF {
		switch p.peek_token.type_ {
		case PlusToken, MinusToken, AsteriskToken, SlashToken:
			e := p.GetNextToken()
			if e != nil {
				return Node{}, e
			}
			node, e = p.parseInfix(node)
			if e != nil {
				return Node{}, e
			}
			break
		default:
			break
		}
	}
	return node, nil
}

func (p *Parser) parsePrefix() (Node, error) {
	switch p.current_token.type_ {
	case MinusToken:
		val := "-" + p.peek_token.literal
		e := p.GetNextToken()
		if e != nil {
			return Node{}, e
		}
		return Node{
			type_: NumberNode,
			val:   val,
			left:  nil,
			right: nil,
		}, e
	case NumberToken:
		val := p.current_token.literal
		return Node{
			type_: NumberNode,
			val:   val,
			left:  nil,
			right: nil,
		}, nil
	default:
		return Node{}, errors.New("unknown token: '" + p.current_token.literal + "'")
	}
}

func (p *Parser) parseInfix(left Node) (Node, error) {
	op := p.current_token

	e := p.GetNextToken()
	if e != nil {
		return Node{}, e
	}

	right, e := p.parseInternal(p.current_token.GetPrecedence())
	if e != nil {
		return Node{}, e
	}

	n := Node{
		type_: OperatorNode,
		val:   op.literal,
		left:  &left,
		right: &right,
	}
	return n, nil
}

func (p *Parser) GetNextToken() error {
	p.current_token = p.peek_token
	n, e := p.lx.GetNextToken()
	if e != nil {
		return e
	}
	p.peek_token = n
	return nil
}
