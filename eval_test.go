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
		_type: NumberNode,
		val:   "10",
		left:  nil,
		right: nil,
	}, want: 10}, {input: Node{
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
	}, want: 12}, {input: Node{
		_type: OperatorNode,
		val:   "*",
		left: &Node{
			_type: NumberNode,
			val:   "3",
		},
		right: &Node{
			_type: OperatorNode,
			val:   "+",
			left: &Node{
				_type: NumberNode,
				val:   "2",
			},
			right: &Node{
				_type: NumberNode,
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
