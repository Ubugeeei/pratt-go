package main

import "testing"

type Tests struct {
	input string
	want  Node
}

func TestParser(t *testing.T) {
	tests :=
		[]Tests{
			{
				input: "10 + 2 * 3",
				want: Node{
					_type: OperatorNode,
					val: Token{
						_type:   Plus,
						literal: "+",
					},
					left: &Node{
						val: Token{
							_type:   Number,
							literal: "10",
						},
					},
					right: &Node{
						_type: OperatorNode,
						val: Token{
							_type:   Asterisk,
							literal: "*",
						},
						left: &Node{
							val: Token{
								_type:   Number,
								literal: "2",
							},
						},
						right: &Node{
							val: Token{
								_type:   Number,
								literal: "3",
							},
						},
					},
				},
			},
		}

	for _, test := range tests {
		t.Run("test", func(t *testing.T) {
			p := NewParser(NewLexer(test.input))
			if got := p.parse(-1); got != test.want {
				t.Errorf("got = %v, want %v", got, test.want)
			}
		})
	}
}
