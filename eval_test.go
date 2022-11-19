package main

import (
	"reflect"
	"testing"
)

type TestCase struct {
	input Node
	want  int
}

func TestEval(t *testing.T) {
	tests := []TestCase{{input: Node{
		type_: NumberNode,
		val:   "10",
		left:  nil,
		right: nil,
	}, want: 10}, {input: Node{
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
	}, want: 12}, {input: Node{
		type_: OperatorNode,
		val:   "*",
		left: &Node{
			type_: NumberNode,
			val:   "3",
		},
		right: &Node{
			type_: OperatorNode,
			val:   "+",
			left: &Node{
				type_: NumberNode,
				val:   "2",
			},
			right: &Node{
				type_: NumberNode,
				val:   "1",
			},
		},
	}, want: 9}}

	for _, test := range tests {
		t.Run("test", func(t *testing.T) {

			if got := eval(test.input); !reflect.DeepEqual(got, test.want) {
				t.Errorf("test failed!")
			}
		})
	}
}
