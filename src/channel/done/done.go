package main

import (
	"fmt"
)

//返回数据只能收数据
func createWorker(id int) worker {
	w := worker{
		make(chan int),
		make(chan bool),
	}
	go doWork(id, w.in, w.done)
	return w

}

func channelDemo() {
	//var c chan int //c == nil

	var workers [10]worker
	for i := 0; i < 10; i++ {
		//channels[i] = make(chan int)
		workers[i] = createWorker(i)
	}

	for i, worker := range workers {
		worker.in <- 'a' + i
		//<-workers[i].done
	}

	for _, worker := range workers {
		<-worker.done
	}

	for i, worker := range workers {
		worker.in <- 'A' + i
		//<-workers[i].done

	}

	//wait for all of them
	for _, worker := range workers {
		<-worker.done
	}

}

type worker struct {
	in   chan int
	done chan bool
}

func doWork(id int, c chan int, done chan bool) {
	//c = make(chan int, 3)
	/*for {
		n, ok := <-c
		if !ok {
			break
		}
		fmt.Printf("Worker %d received %d \n", id, n)
	}*/

	//也可以使用range的方式替代上面
	for n := range c {
		fmt.Printf("Worker %d received %c \n", id, n)
		done <- true

	}

}

func main() {
	channelDemo()
}
