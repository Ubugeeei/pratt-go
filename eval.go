package main

import (
	"strconv"
)

func eval(node Node) float64 {
	switch node.type_ {
	case NumberNode:
		{
			n, _ := strconv.ParseFloat(node.val, 64)
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
		}
	}

	return 0 // TODO: error handling
}
