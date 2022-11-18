package main

import "testing"

func TestLexer(t *testing.T) {
	lx := NewLexer("10+2")

	tests := []Token{
		{
			_type:   Number,
			literal: "10",
		},
		{
			_type:   Plus,
			literal: "+",
		},
		{
			_type:   Number,
			literal: "2",
		},
	}

	for _, want := range tests {
		t.Run("test", func(t *testing.T) {
			if got := lx.nextToken(); got != want {
				t.Errorf("got = %v, want %v", got, want)
			}
		})
	}
}
