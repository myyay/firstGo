package main

import (
	"fmt"
	"sync"
)

//返回数据只能收数据
func createWorker2(id int, wg *sync.WaitGroup) worker2 {
	w := worker2{
		in: make(chan int),
		//把done的方法自定义 抽象出来
		done: func() {
			wg.Done()
		},
	}
	go doWork2(id, w)
	return w

}

func channelDemo2() {
	//var c chan int //c == nil

	var workers [10]worker2
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		//channels[i] = make(chan int)
		workers[i] = createWorker2(i, &wg)
	}

	//CountDownLatch
	wg.Add(20)
	for i, worker := range workers {
		worker.in <- 'a' + i
		//<-workers[i].done
	}

	for i, worker := range workers {
		worker.in <- 'A' + i
		//<-workers[i].done

	}

	wg.Wait()
}

type worker2 struct {
	in   chan int
	done func()
}

func doWork2(id int, w worker2) {
	//c = make(chan int, 3)
	/*for {
		n, ok := <-c
		if !ok {
			break
		}
		fmt.Printf("Worker %d received %d \n", id, n)
	}*/

	//也可以使用range的方式替代上面
	for n := range w.in {
		fmt.Printf("Worker %d received %c \n", id, n)
		w.done()

	}

}

func main() {
	channelDemo2()
}
