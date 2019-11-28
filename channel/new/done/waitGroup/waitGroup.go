package main

//怎么知道worker都执行完成了呢,加一个标识字段 done

//还有一种方式叫waitGroup

import (
	"fmt"
	"sync"
)

func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go doWork(id, w)
	return w
}

type worker struct {
	in chan int
	//通过done方法把waitGroup就给抽象掉了，可以在createWorker里自己定义要做什么
	done func()
}

func doWork(id int, w worker) {

	for n := range w.in {
		fmt.Printf("Worker %d received %c \n", id, n)
		w.done()
	}
}

func channelDemo() {
	var workers [10]worker
	var wg sync.WaitGroup
	wg.Add(20)

	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}

	for i, worker := range workers {
		worker.in <- 'a' + i
	}

	for i, worker := range workers {
		worker.in <- 'A' + i
	}

	wg.Wait()

}

func main() {
	channelDemo()

}
