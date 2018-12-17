package main

import (
	"ccmouse/dataBasic/tree"
	"fmt"
)

func main() {
	var root tree.Node

	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left.SetValue(4)

	fmt.Print("In-oder traversal:")
	root.Traverse()

	c := root.TraverseWithChannel()
	maxNodeValue := -1
	for node := range c {
		if node.Value >= maxNodeValue {
			maxNodeValue = node.Value
		}
	}
	fmt.Println("max node value is:", maxNodeValue)
}
