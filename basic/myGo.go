package main

import "fmt"
import "math/cmplx"
import "math"

//包内部变更  并非全局变量

var (
	aa = 3
	ss = "kkk"
	bb = true
)

func variablesZeroValue() {
	var a int
	var s string
	fmt.Printf("%d , %q \n", a, s)

}

func variablesInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

func variableShorter() {
	//这种赋值只能在函数内使用
	a, b, c, s := 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

func euler() {
	c := 3 + 4i
	fmt.Println("euler:", cmplx.Abs(c))
}

func triangle() {
	var a, b int = 3, 4

	fmt.Println("triangle:", calcTriangle(a, b))
}

func calcTriangle(a, b int) int {
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	return c
}

func enums() {
	const (
		cpp    = 1
		java   = 2
		python = 3
		golang = 4
	)
	fmt.Println(cpp, java, python, golang)

	const (
		c = iota
		_
		j
		p
		g
	)
	fmt.Println("c, j, p ,g = ", c, j, p, g)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println("b, kb, mb, gb, tb, pb", b, kb, mb, gb, tb, pb)

}

func main() {
	fmt.Println("Hello World ")
	variablesZeroValue()
	variablesInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, bb, ss)
	euler()
	triangle()

	enums()
	// string bool
	//uint int int8 int16 int32
	//rune -> char类型 是32字节
	//float32 float64 complect64 complect128 (复数实部和虚部各64位)
}
