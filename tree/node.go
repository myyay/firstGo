package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

//由于没有构造函数，工厂函数代替构造，可以自定义一些操作
func CreateNode(value int) *Node {
	//go lang允许返回局部变量的地址给外部 C++不行
	return &Node{Value: value}

}

/** 前置参数表示是结构体内的参数   也是传值的 */
func (node *Node) Print() {
	fmt.Print(node.Value, " ")
}

/** setValue不会影响原值  要想修改原值 需要传指针 */
func (node *Node) SetValue(value int) {
	if nil == node {
		fmt.Println("Setting value to nil. Ignore")
		return
	}
	node.Value = value
}
