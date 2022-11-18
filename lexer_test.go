package main

import "testing"

func TestLexer(t *testing.T) {
	lx := NewLexer("1 + 2")

	tests := []Token{
		{
			_type:   Number,
			literal: "1",
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

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := lx.nextToken(); got != tt {
				t.Errorf("add() = %v, want %v", got, tt)
			}
		})
	}
}
