package main

import (
	"fmt"
)

type Node struct {
	data  string
	left  *Node
	right *Node
}

func NewNode(data string) *Node {
	return &Node{
		data:  data,
		left:  nil,
		right: nil,
	}
}

func (root *Node) insert(data string) *Node {
	if &root == nil {
		return NewNode(data)
	}
	if data < root.data {
		root.left = root.left.insert(data)
	} else {
		root.right = root.right.insert(data)
	}
	return root
}

func (root *Node) search(key string) *Node {
	if &root == nil || root.data == key {
		return root
	}
	if key < root.data {
		return root.left.search(key)
	}
	return root.right.search(key)
}

func inOrderTraversal(root *Node) {
	if root != nil {
		inOrderTraversal(root.left)
	}
	fmt.Println(root.data)
	inOrderTraversal(root.right)
}
