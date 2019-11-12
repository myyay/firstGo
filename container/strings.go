package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Hello， 你好，世界！"
	//获取字节数
	fmt.Println(len(s))
	fmt.Printf("%X\n", []byte(s))

	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)

	}

	fmt.Println()

	for i, ch := range s {
		fmt.Printf("(%d,%X) ", i, ch) //ch is rune

	}

	fmt.Println()

	//获取字符数
	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}

	fmt.Println()
	//会开一个rune的slice给我们
	for i, ch := range []rune(s) {
		fmt.Printf("(%d,%c)", i, ch)
	}
	fmt.Println()

}
