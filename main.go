package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Usage: go run main.go [-i] [expression]")
		os.Exit(1)
	}

	if isInteractive(os.Args) {
		runInteractive()
	} else {
		run(os.Args[1])
	}
}

func run(expr string) {
	lexer := NewLexer(expr)
	parser := NewParser(lexer)

	ast, e := parser.Parse()
	if e != nil {
		fmt.Println(e)
		return
	}

	res, e := eval(ast)
	if e != nil {
		fmt.Println(e)
		return
	}

	fmt.Println(res)
}

func runInteractive() {
	fmt.Println("input your expression")
	for {
		fmt.Print("> ")

		var input string
		fmt.Scanln(&input)

		if input == "exit" {
			fmt.Println("Bye!")
			break
		}

		lexer := NewLexer(input)
		parser := NewParser(lexer)

		ast, e := parser.Parse()
		if e != nil {
			fmt.Println(e)
			continue
		}

		res, e := eval(ast)
		if e != nil {
			fmt.Println(e)
			continue
		}

		fmt.Println(res)
	}
}

func isInteractive(argv []string) bool {
	for _, arg := range argv {
		if arg == "-i" {
			return true
		}
	}
	return false
}
