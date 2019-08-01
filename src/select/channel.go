package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c

}

func worker(id int, c chan int) {
	for n := range c {
		time.Sleep(time.Second)
		fmt.Printf("Worker %d received %d \n", id, n)

	}

}

func main() {

	var c1, c2 = generator(), generator()

	worker := createWorker(0)

	/*for {
		select {
		case n := <-c1:
			fmt.Println("Received from c1 :", n)
		case n := <-c2:
			fmt.Println("Received from c2 :", n)

			//default:
			//fmt.Println("No value received")

		}
	}*/

	n := 0
	/*hasValue := false


	for {
		//必须在for循环里赋值 nil会自动被select 忽略
		var activeWorker chan<- int
		if hasValue {
			activeWorker = worker
		}
		select {
		case n = <-c1:
			hasValue = true
		case n = <-c2:
			hasValue = true
		case activeWorker <- n:
			hasValue = false

			//default:
			//fmt.Println("No value received")

		}
	}*/

	//增加队列 如果处理时间过长 会被丢弃掉

	var values []int
	tm := time.After(10 * time.Second)

	tick := time.Tick(time.Second)
	for {
		//必须在for循环里赋值 nil会自动被select 忽略
		var activeWorker chan<- int
		var activeValue int
		if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]
		}
		select {
		case n = <-c1:
			values = append(values, n)
		case n = <-c2:
			values = append(values, n)
		case activeWorker <- activeValue:
			values = values[1:]

		//如果再次生成数据之间超过800毫秒 就会打印
		case <-time.After(800 * time.Millisecond):
			fmt.Println("timeout")
		//每秒监控
		case <-tick:
			fmt.Println("queue length:", len(values))
		case <-tm:
			fmt.Println("Bye")
			return

			//default:
			//fmt.Println("No value received")

		}
	}
}
