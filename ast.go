package main

type Node struct {
	_type NodeType
	val   Token
	left  *Node
	right *Node
}

func (n Node) toString() string {
	var left string
	if n.left == nil {
		left = "nil"
	} else {
		left = n.left.toString()
	}

	var right string
	if n.right == nil {
		right = "nil"
	} else {
		right = n.right.toString()
	}

	return "{ val: " + n.val.literal + ", left: " + left + ", right: " + right + " }"
}

type NodeType int

const (
	NumberNode NodeType = iota
	OperatorNode
)
