package main

import "testing"

func TestLexer(t *testing.T) {
	lx := NewLexer("10     + 2 / 3")

	tests := []Token{
		{
			type_:   NumberToken,
			literal: "10",
		},
		{
			type_:   PlusToken,
			literal: "+",
		},
		{
			type_:   NumberToken,
			literal: "2",
		},
		{
			type_:   SlashToken,
			literal: "/",
		},
		{
			type_:   NumberToken,
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
