package main

import (
	"fmt"
	"runtime"
	"time"
)

//执行命令 go run -race goroutine2.go
//这个跟goroutine 的区别是 不把i传进到函数里
//会报 runtime error: index out of range [10] with length 10
func main() {
	var a [10]int
	for i := 0; i < 10; i++ {
		//当i 加到10的时候 进不了循环 ，但routine依然在运行 导致数据越界
		go func() { //race condition 数据访问冲突
			for {
				//fmt.Printf("Hello from goroutine %d \n", i)
				//如果没有IO操作 会死在这
				a[i]++
				//手动交出控制权
				runtime.Gosched()
			}
		}()
	}
	time.Sleep(time.Millisecond)
	fmt.Println("a: ", a)
}
