package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

func CreateNode(value int) *Node {
	//局部变量的地址
	return &Node{Value: value}

}

/** 前置参数表示是结构体内的参数   也是传值的 */
func (node *Node) Print() {
	fmt.Print(node.Value, " ")
}

/** setValue不会影响原值  要想修改原值 需要传指针 */
func (node *Node) SetValue(value int) {
	node.Value = value
}
