package tree

import "fmt"

func (node *Node) Traverse() {

	node.TraverseFunc(func(node *Node) {
		node.Print()
	})

	fmt.Println()
}

//把函数做成参数
func (node *Node) TraverseFunc(f func(node *Node)) {

	if node == nil {
		return
	}
	//如果left是空也可以使用
	node.Left.Traverse()
	f(node)
	node.Right.Traverse()
}
