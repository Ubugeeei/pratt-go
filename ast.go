package main

type Node struct {
	type_ NodeType
	val   string
	left  *Node
	right *Node
}

type NodeType int

const (
	NumberNode NodeType = iota
	OperatorNode
)

func (n Node) ToString() string {
	var left string
	if n.left == nil {
		left = "nil"
	} else {
		left = n.left.ToString()
	}

	var right string
	if n.right == nil {
		right = "nil"
	} else {
		right = n.right.ToString()
	}

	return "{ val: " + n.val + ", left: " + left + ", right: " + right + " }"
}
