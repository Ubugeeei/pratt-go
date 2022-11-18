package main

type Node struct {
	_type NodeType
	val   Token
	left  *Node
	right *Node
}

type NodeType int

const (
	NumberNode NodeType = iota
	OperatorNode
)
