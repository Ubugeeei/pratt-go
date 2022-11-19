package main

import (
	"fmt"
	"os"
)

func main() {
	argv := os.Args

	var isInteractive bool
	for _, arg := range argv {
		if arg == "-i" {
			isInteractive = true
		}
	}

	if isInteractive {
		fmt.Println("input your expression")
		for {
			fmt.Print("> ")

			var input string
			fmt.Scanln(&input)

			if input == "exit" {
				fmt.Println("Bye!")
				return
			}

			lexer := NewLexer(input)
			parser := NewParser(lexer)

			ast := parser.parse(0)
			res, e := eval(ast)
			if e != nil {
				fmt.Println(e)
			} else {
				fmt.Println(res)
			}
		}
	} else {
		expr := argv[1]
		lexer := NewLexer(expr)
		parser := NewParser(lexer)

		ast := parser.parse(0)
		res, e := eval(ast)
		if e != nil {
			fmt.Println(e)
		} else {
			fmt.Println(res)
		}
	}
}
