package main

import (
	"bufio"
	"errors"
	"firstGo/functional/fibonacci/fib"
	"fmt"
	"os"
)

func tryDefer() {
	//表示函数退出之前打印
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	panic("error occurred")
	fmt.Println(4)

}

func writeFile(filename string) {
	file, err := os.Create(filename)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fib.Fibonacci()

	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}

}

func tryErr(filename string) {
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)

	err = errors.New("this is a custom error")

	if pathError, ok := err.(*os.PathError); !ok {
		panic(err)
	} else {
		fmt.Printf("%s, %s, %s", pathError.Op, pathError.Path, pathError.Err)
		fmt.Println()
	}

	if err != nil {
		panic(err)
	}

	defer file.Close()

}

func main() {
	//tryDefer()
	//writeFile("fib.txt")

	tryErr("fib.txt")

}
