package main

type Node struct {
	_type NodeType
	val   string
	left  *Node
	right *Node
}

type NodeType int

const (
	NumberNode NodeType = iota
	OperatorNode
)

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

	return "{ val: " + n.val + ", left: " + left + ", right: " + right + " }"
}
