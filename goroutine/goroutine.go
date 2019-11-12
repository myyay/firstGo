package main

import (
	"fmt"
	"runtime"
	"time"
)

func printHello() {

}

func main() {
	var a [10]int
	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				//fmt.Printf("Hello from goroutine %d \n", i)
				//如果没有IO操作 会死在这
				a[i]++
				//手动交出控制权
				runtime.Gosched()
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
	fmt.Println("a: ", a)
}
