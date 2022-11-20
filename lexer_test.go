package main

import (
	"reflect"
	"testing"
)

type LexerTestCase struct {
	input string
	want  []Token
}

func TestLexer(t *testing.T) {
	tests := []LexerTestCase{
		{
			input: "10",
			want: []Token{
				{
					type_:   NumberToken,
					literal: "10",
				},
				{
					type_:   EOF,
					literal: "",
				},
			},
		},
		{
			input: "10 + 2",
			want: []Token{
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
					type_:   EOF,
					literal: "",
				},
			},
		},
		{
			input: "10     + 2 / 3",
			want: []Token{
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
				{
					type_:   EOF,
					literal: "",
				},
			},
		},
	}

	for _, testCase := range tests {
		t.Run("test", func(t *testing.T) {
			lx := NewLexer(testCase.input)
			var got []Token
			for {
				tok, _ := lx.GetNextToken()
				got = append(got, tok)
				if tok.type_ == EOF {
					break
				}
			}
			if !reflect.DeepEqual(got, testCase.want) {
				t.Errorf("got = %v, want %v", got, testCase.want)
			}
		})
	}
}
