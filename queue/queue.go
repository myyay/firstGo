package queue

import (
	"fmt"
)

//interface{} 表示是任何类型
type Queue []interface{}

func (q *Queue) Push(v interface{}) {
	*q = append(*q, v)

}

func (q *Queue) Pop() interface{} {

	if q.IsEmpty() {
		panic(fmt.Sprint("列表为空，不能pop"))
	}

	head := (*q)[0]
	*q = (*q)[1:]
	return head

}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0

}
