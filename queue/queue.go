package queue

import (
	"fmt"
)

//interface{} 表示是任何类型
type Queue []interface{}

// Pushes the element into the queue.
// 		e.g. q.Push(123)
func (q *Queue) Push(v interface{}) {
	*q = append(*q, v)

}

// Pops element from head.
func (q *Queue) Pop() interface{} {

	if q.IsEmpty() {
		panic(fmt.Sprint("列表为空，不能pop"))
	}

	head := (*q)[0]
	*q = (*q)[1:]
	return head

}

// Returns if the queue is empty or not.
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0

}

//使用 go doc 命令来查看文档

//执行  go doc
// package queue // import "firstGo/queue"
//
// type Queue []interface{}

// go doc Queue 查看方法信息

//go doc IsEmpty 查看方法信息

//go doc -help 查看帮助信息

//godoc -http=localhost:6060

//注释格式可以到文档里
