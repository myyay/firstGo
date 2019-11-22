package tree

import "fmt"

//中序遍历
func (node *Node) NormalTraverse() {
	if nil == node {
		return
	}

	node.Left.NormalTraverse()
	node.Print()
	node.Right.NormalTraverse()

}

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

func (node *Node) TraverseWithChannel() chan *Node {
	out := make(chan *Node)
	go func() {
		node.TraverseFunc(func(innerNode *Node) {
			out <- innerNode
		})
		close(out)
	}()
	return out
}
