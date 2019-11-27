package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var a [10]int
	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				//fmt.Printf("Hello from goroutine %d \n", i)
				//如果没有IO操作 会死在这 goroutine不碰到io操作，不会交出控制权
				a[i]++
				//手动交出控制权
				runtime.Gosched()
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
	fmt.Println("a: ", a)

	//goroutine 的切换点
	//I/O select
	//channel
	//等待锁
	//函数调用(有时)
	//runtime.Gosched()
}
