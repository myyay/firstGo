package main

import (
	"firstGo/src/queue"
	"fmt"
)

func main() {

	q := queue.Queue{1}

	fmt.Println(q.IsEmpty())

	q.Push(2)
	q.Push(3)
	q.Push(4)

	fmt.Println(q)

	fmt.Println(q.Pop(), q.Pop(), q.Pop())

	fmt.Println(q.IsEmpty())

	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())

	q.Push("abc")
	fmt.Println(q.Pop())

}
