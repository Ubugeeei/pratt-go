package main

import (
	"reflect"
	"testing"
)

type Tests struct {
	input string
	want  Node
}

func TestParser(t *testing.T) {
	tests :=
		[]Tests{
			{
				input: "10",
				want: Node{
					_type: NumberNode,
					val:   "10",
					left:  nil,
					right: nil,
				},
			},
			{
				input: "-10",
				want: Node{
					_type: NumberNode,
					val:   "-10",
					left:  nil,
					right: nil,
				},
			},
			{
				input: "10 + 2",
				want: Node{
					_type: OperatorNode,
					val:   "+",
					left: &Node{
						_type: NumberNode,
						val:   "10",
					},
					right: &Node{
						_type: NumberNode,
						val:   "2",
					},
				},
			},
			{
				input: "10 + 2 * 3",
				want: Node{
					_type: OperatorNode,
					val:   "+",
					left: &Node{
						_type: NumberNode,
						val:   "10",
					},
					right: &Node{
						_type: OperatorNode,
						val:   "*",
						left: &Node{
							_type: NumberNode,
							val:   "2",
						},
						right: &Node{
							_type: NumberNode,
							val:   "3",
						},
					},
				},
			},
		}

	for _, test := range tests {
		t.Run("test", func(t *testing.T) {
			p := NewParser(NewLexer(test.input))
			got := p.parse(-1)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("test failed!\ninput:" + test.input + "\ngot: " + got.toString())
			}
		})
	}
}
