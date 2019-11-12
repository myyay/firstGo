package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"unsafe"
)

var (
	aa = 3
	ss = "kkk"
	bb = true
)

func variableZeroValue() {
	var a int
	var s string
	// 这里用了%q
	fmt.Printf("%d %q \n", a, s)
}

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "abc"
	fmt.Println(a, b, c, s)
}

func variableShorter() {
	var a, b, c, s = 3, 4, true, "abc"
	b = 5
	fmt.Println(a, b, c, s)
}

func euler() {
	fmt.Printf("%.3f\n", cmplx.Exp(1i*math.Pi)+1)
}

func triangle() {
	var a, b = 3, 4
	fmt.Println(calcTriangle(a, b))
}

func calcTriangle(a, b int) int {
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	return c
}

func consts() {
	const (
		filename = "abc.txt"
		a, b     = 3, 4
	)
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(c)
}

func enums() {
	const (
		cpp = iota
		_
		python
		golang
		javascript
	)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(cpp, javascript, python, golang)
	fmt.Println(b, kb, mb, gb, tb, pb)

}

func sizeOfType() {
	// string bool
	//uint int int8 int16 int32
	//rune -> char类型 是32字节
	//float32 float64 complect64 complect128 (复数实部和虚部各64位)
	var iu uint = 1
	var i int = 2
	var i8 int8 = 3
	var i16 int16 = 4
	var i32 int32 = 5

	fmt.Println("int占用: ", unsafe.Sizeof(iu), unsafe.Sizeof(i), unsafe.Sizeof(i8), unsafe.Sizeof(i16), unsafe.Sizeof(i32))
	var s string = "abc"
	var b bool = true

	fmt.Println("string/boo占用: ", unsafe.Sizeof(s), unsafe.Sizeof(b))

}

func main() {
	fmt.Println(1i)
	fmt.Println(aa, bb, ss)
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	euler()
	triangle()
	consts()
	enums()
	sizeOfType()
}
