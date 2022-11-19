package main

import (
	"strconv"
)

func eval(node Node) int {
	switch node.type_ {
	case NumberNode:
		{
			n, _ := strconv.Atoi(node.val)
			return n
		}
	case OperatorNode:
		switch node.val {
		case "+":
			return eval(*node.left) + eval(*node.right)
		case "-":
			return eval(*node.left) - eval(*node.right)
		case "*":
			return eval(*node.left) * eval(*node.right)
		case "/":
			return eval(*node.left) / eval(*node.right)
		case "%":
			return eval(*node.left) % eval(*node.right)
		}
	}

	return 0 // TODO: error handling
}
