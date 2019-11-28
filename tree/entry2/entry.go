package main

import (
	"firstGo/tree"
	"fmt"
)

func main() {
	var root tree.Node
	root = tree.Node{Value: 3}
	//golang里的指针类型也可以用.的方法往下操作。 而其他语言是箭头->
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left.SetValue(4)

	fmt.Println("In-order traversal: ")
	//0 2 3 4 5
	root.NormalTraverse()

	fmt.Println()

	fmt.Println("traverse by traverse func")
	root.Traverse()
	nodeCount := 0
	root.TraverseFunc(func(node *tree.Node) {
		nodeCount++
	})

	fmt.Println("Node count: ", nodeCount)

	out := root.TraverseWithChannel()
	maxNodeValue := 0
	for n := range out {
		if n.Value > maxNodeValue {
			maxNodeValue = n.Value
		}
		n.Print()
	}

	fmt.Println()
	fmt.Printf("maxValue of Tree is %d \n", maxNodeValue)

}
