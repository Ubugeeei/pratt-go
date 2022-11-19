package main

import (
	"errors"
	"strconv"
)

func eval(node Node) (float64, error) {
	switch node.type_ {
	case NumberNode:
		{
			n, e := strconv.ParseFloat(node.val, 64)
			return n, e
		}
	case OperatorNode:
		switch node.val {
		case "+":
			{
				left, errL := eval(*node.left)
				right, errR := eval(*node.right)
				if errL != nil {
					return 0, errL
				}
				if errR != nil {
					return 0, errR
				}
				return left + right, nil
			}
		case "-":
			{
				left, errL := eval(*node.left)
				right, errR := eval(*node.right)
				if errL != nil {
					return 0, errL
				}
				if errR != nil {
					return 0, errR
				}
				return left - right, nil
			}
		case "*":
			{
				left, errL := eval(*node.left)
				right, errR := eval(*node.right)
				if errL != nil {
					return 0, errL
				}
				if errR != nil {
					return 0, errR
				}
				return left * right, nil
			}
		case "/":
			{
				left, errL := eval(*node.left)
				right, errR := eval(*node.right)
				if errL != nil {
					return 0, errL
				}
				if errR != nil {
					return 0, errR
				}
				return left / right, nil
			}
		}
	}

	return 0, errors.New("unreachable")
}
