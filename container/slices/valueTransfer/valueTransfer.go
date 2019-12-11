package main

import (
	"fmt"
)

type myType struct {
	String string
	Int    int
}

func modifyMyType(t myType) {
	fmt.Printf("myType position: %p, value: %v \n", &t, t)

	t.String = "a"
	t.Int = 1

}

func modifyMap(m map[string]int) {
	fmt.Printf("map position: %p, value: %v \n", m, m)

	m["a"] = 1
	m["b"] = 2
	m["c"] = 3
}

func modifySlice(s []int) {
	fmt.Printf("slice position: %p, value: %v \n", s, s)

	s[0] = 1
	s = append(s, 888)
}
func modifyArray(arr [5]int) {
	fmt.Printf("array position: %p, value: %v \n", &arr, arr)

	arr[0] = 1
}

func main() {
	//map 在创建的的时候，返回的其实是指针  chan类型和map一样！
	//添加新元素是有效的，因为是指针
	m1 := map[string]int{
		"a": 8,
		"b": 9,
	}

	m2 := make(map[string]int, 6)

	m2["a"] = 5
	m2["b"] = 6

	fmt.Printf("map: m1 position: %p, value: %v \n", m1, m1)
	modifyMap(m1)
	fmt.Printf("map: after modify m1 position: %p, value: %v \n", m1, m1)

	fmt.Printf("map: m2 position: %p, value: %v \n", m2, m2)
	modifyMap(m2)
	fmt.Printf("map: after modify m2 position: %p, value: %v \n", m2, m2)

	fmt.Println("===========================")
	//普通对象肯定是值拷贝
	mt := myType{
		String: "str",
		Int:    333,
	}

	fmt.Printf("my type: mt position: %p, value: %v \n", &mt, mt)
	modifyMyType(mt)
	fmt.Printf("my type: after modify mt position: %p, value: %v \n", &mt, mt)

	fmt.Println("===========================")

	//slice 相当于以下结构体，只能改变内部变量的值，不能添加元素删除元素。如果要添加元素或其他操作，要传地址作为参数！！
	//type slice struct {
	//	array unsafe.Pointer
	//	len int
	//	cap int
	//}

	s := []int{2, 3, 6, 7, 9}

	fmt.Printf("slice : s position: %p, value: %v \n", s, s)
	modifySlice(s)
	fmt.Printf("slice : after modify s position: %p, value: %v \n", s, s)

	fmt.Println("===========================")

	//数组在传递时也是值拷贝
	arr := [...]int{2, 3, 6, 7, 9}

	fmt.Printf("array : s position: %p, value: %v \n", &arr, arr)
	modifyArray(arr)
	fmt.Printf("array : after modify s position: %p, value: %v \n", &arr, arr)

}
