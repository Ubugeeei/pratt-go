package main

type Parser struct {
	lx            *Lexer
	current_token Token
	peek_token    Token
}

func NewParser(lx *Lexer) *Parser {
	p := new(Parser)
	p.lx = lx
	p.current_token = lx.nextToken()
	p.peek_token = lx.nextToken()
	return p
}

// pratt parser
func (p *Parser) parse(precedence int) Node {
	node := p.parsePrefix()

	for precedence < p.peek_token.getPrecedence() && p.peek_token._type != EOF {
		switch p.peek_token._type {
		case Plus, Minus, Asterisk, Slash, Percent:
			p.nextToken()
			node = p.parseInfix(node)
			break
		default:
			break
		}
	}
	return node
}

func (p *Parser) parsePrefix() Node {
	switch p.current_token._type {
	case Minus:
		t := Token{
			_type:   Number,
			literal: "-" + p.peek_token.literal,
		}
		return Node{
			_type: NumberNode,
			val:   t,
			left:  nil,
			right: nil,
		}
	default:
		t := p.current_token
		return Node{
			_type: NumberNode,
			val:   t,
			left:  nil,
			right: nil,
		}
	}
}

func (p *Parser) parseInfix(left Node) Node {
	op := p.current_token

	p.nextToken()
	right := p.parse(p.current_token.getPrecedence())

	n := Node{
		_type: OperatorNode,
		val:   op,
		left:  &left,
		right: &right,
	}
	return n
}

func (p *Parser) nextToken() {
	p.current_token = p.peek_token
	p.peek_token = p.lx.nextToken()
}
