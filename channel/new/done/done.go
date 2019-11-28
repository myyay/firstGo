package main

//怎么知道worker都执行完成了呢,加一个标识字段 done

//还有一种方式叫waitGroup

import (
	"fmt"
)

func createWorker(id int) worker {
	w := worker{
		in:   make(chan int),
		done: make(chan bool),
	}
	go doWork(id, w.in, w.done)
	return w
}

type worker struct {
	in   chan int
	done chan bool
}

func doWork(id int, c chan int, done chan bool) {

	for n := range c {
		fmt.Printf("Worker %d received %c \n", id, n)
		go func() { done <- true }()
	}
}

func channelDemo() {
	var workers [10]worker

	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i)
	}

	for i, worker := range workers {
		worker.in <- 'a' + i
	}

	for _, worker := range workers {
		<-worker.done

	}

	for i, worker := range workers {
		worker.in <- 'A' + i
	}

	for _, worker := range workers {
		<-worker.done
	}

}

func main() {
	channelDemo()

}
