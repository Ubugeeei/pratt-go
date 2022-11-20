package main

import (
	"reflect"
	"testing"
)

type ParserTestCase struct {
	input string
	want  Node
}

func TestParser(t *testing.T) {
	tests :=
		[]ParserTestCase{
			{
				input: "10",
				want: Node{
					type_: NumberNode,
					val:   "10",
					left:  nil,
					right: nil,
				},
			},
			{
				input: "-10",
				want: Node{
					type_: NumberNode,
					val:   "-10",
					left:  nil,
					right: nil,
				},
			},
			{
				input: "10 + 2",
				want: Node{
					type_: OperatorNode,
					val:   "+",
					left: &Node{
						type_: NumberNode,
						val:   "10",
					},
					right: &Node{
						type_: NumberNode,
						val:   "2",
					},
				},
			},
			{
				input: "10 + 2 * 3",
				want: Node{
					type_: OperatorNode,
					val:   "+",
					left: &Node{
						type_: NumberNode,
						val:   "10",
					},
					right: &Node{
						type_: OperatorNode,
						val:   "*",
						left: &Node{
							type_: NumberNode,
							val:   "2",
						},
						right: &Node{
							type_: NumberNode,
							val:   "3",
						},
					},
				},
			},
		}

	for _, test := range tests {
		t.Run("test", func(t *testing.T) {
			p := NewParser(NewLexer(test.input))
			got, e := p.Parse()
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("test failed!\ninput:" + test.input + "\ngot: " + got.ToString())
			} else if e != nil {
				t.Errorf("%d", e)
			}
		})
	}
}
