package main

import (
	"firstGo/tree"
	"fmt"
)

type myTreeNode struct {
	node *tree.Node
}

/** 前序遍历 */
func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}

	left := myTreeNode{myNode.node.Left}
	right := myTreeNode{myNode.node.Right}

	left.postOrder()
	right.postOrder()
	myNode.node.Print()

}

func main() {

	var root tree.Node

	root = tree.Node{Value: 3}

	root.Left = &tree.Node{}

	root.Right = &tree.Node{5, nil, nil}

	// 指针也可以用.
	root.Right.Left = new(tree.Node)
	root.Right.Right = tree.CreateNode(2)

	fmt.Println(root)

	root.Print()

	root.Right.Left.SetValue(4)

	root.Right.Left.Print()

	pRoot := &root

	pRoot.Print()
	pRoot.SetValue(200)

	pRoot.Print()

	root.Print()

	fmt.Println()

	root.Traverse()

	fmt.Println("==========================")

	myRoot := myTreeNode{&root}

	myRoot.postOrder()

	fmt.Println()

	fmt.Println("==========================")
	nodeCount := 0
	root.TraverseFunc(func(node *tree.Node) {
		nodeCount++
	})

	fmt.Println("nodeCount:", nodeCount)

	fmt.Println("-------------------------------")

	c := root.TraverseWithChannel()

	maxNodeValue := 0
	for node := range c {

		if node.Value > maxNodeValue {
			maxNodeValue = node.Value
		}

	}
	fmt.Println("Max node value:", maxNodeValue)

}
