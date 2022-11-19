package main

import "testing"

func TestLexer(t *testing.T) {
	lx := NewLexer("10     + 2 / 3")

	tests := []Token{
		{
			type_:   Number,
			literal: "10",
		},
		{
			type_:   Plus,
			literal: "+",
		},
		{
			type_:   Number,
			literal: "2",
		},
		{
			type_:   Slash,
			literal: "/",
		},
		{
			type_:   Number,
			literal: "3",
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
