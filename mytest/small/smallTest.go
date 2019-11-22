package main

import (
	"fmt"
	"strings"
)

func main() {
	//var a = rand.Int()
	a := 123
	fmt.Println(a)
	fmt.Printf("%12d\n", a)
	s := fmt.Sprintf("%12d\n", a)
	news := strings.Replace(s[len(s)-13:], " ", "0", -1)
	fmt.Println(news)
}
