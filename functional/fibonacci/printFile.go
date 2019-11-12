package main

import (
	"bufio"
	"firstGo/functional/fibonacci/fib"
	"fmt"
	"io"
	"strings"
)

/*func Fibonacci() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}*/
//可以给函数实现接口
type intGen func() int

func (g intGen) Read(p []byte) (n int, err error) {
	//直接调一下
	next := g()
	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)

}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}

func main() {
	var f intGen = fib.Fibonacci()
	printFileContents(f)

}
