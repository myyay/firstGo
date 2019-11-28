package main

import (
	"fmt"
	"time"
)

//返回数据只能收数据
func createWorker(id int) chan<- int {
	c := make(chan int)
	go func() {
		for {
			fmt.Printf("Worker %d received %c \n", id, <-c)
		}
	}()
	return c

}

func channelDemo() {
	//var c chan int //c == nil

	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		//channels[i] = make(chan int)
		channels[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}

	//c <- 1
	//c <- 2
	/*n := <-c
	fmt.Println(n)*/
	time.Sleep(time.Millisecond)

}

func worker(id int, c chan int) {
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
		fmt.Printf("Worker %d received %d \n", id, n)

	}
}

func bufferedChannel() {
	//增加缓冲区
	c := make(chan int, 3)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	time.Sleep(time.Millisecond)
}

func channelClose() {
	c := make(chan int)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	c <- 'A'
	//close 一直会收到空数据
	close(c)
	time.Sleep(time.Millisecond)
}

func main() {
	//channelDemo()
	//bufferedChannel()
	channelClose()
}
