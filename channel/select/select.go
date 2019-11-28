package main

import (
	"fmt"
	"math/rand"
	"time"
)

func worker(id int, c chan int) {
	for n := range c {
		fmt.Printf("Worker %d received %d \n", id, n)
		time.Sleep(time.Duration(time.Millisecond * 1000))
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go func() {
		worker(id, c)
	}()
	return c
}

func generate() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			//sleep 1500ms以内的数字
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1500)))
			out <- i
			i++
		}

	}()
	return out
}

func main() {

	var c1, c2 = generate(), generate()
	//w := createWorker(0)
	worker := createWorker(0)
	/*
		//有问题  n必须收下来 然后排队。 否则如果生成的太快，消费太慢，就会丢掉中间数据
		n := 0
		hasValue := false
		for {
			//必须在for循环里赋值 nil会自动被select 忽略
			var activeWorker chan<- int
			if hasValue {
				activeWorker = worker
			}
			select {
			//这里是收
			case n = <-c1:
				//fmt.Println("received from c1: ", n)
				hasValue = true
			case n = <-c2:
				//fmt.Println("received from c2: ", n)
				hasValue = true
			//这里是发 这样就变成又可以收 又可以发
			case activeWorker <- n:
				hasValue = false

				//default:
				//fmt.Println("no value received")

			}
		}
	*/

	//想让跑10s后退出
	tm := time.After(time.Second * 10)
	tick := time.Tick(time.Second)

	var values []int
	for {
		//必须在for循环里赋值 nil会自动被select 忽略
		var activeWorker chan<- int
		var activeValue int
		if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]

		}
		select {
		//这里是收
		case n := <-c1:
			//fmt.Println("received from c1: ", n)
			values = append(values, n)
		case n := <-c2:
			//fmt.Println("received from c2: ", n)
			values = append(values, n)
		//这里是发 这样就变成又可以收 又可以发
		case activeWorker <- activeValue:
			values = values[1:]
		case <-time.After(time.Duration(time.Millisecond * 800)):
			fmt.Println("timeout")
		case <-tick:
			fmt.Println("queue len : ", len(values))
		case <-tm:
			fmt.Println("bye")
			return
			//default:
			//fmt.Println("no value received")

		}
	}

}
