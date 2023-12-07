package main

import (
	"container/list"
	"fmt"
)

type Queue struct {
	v *list.List
}

func (q *Queue) Push(val interface{}) {
	q.v.PushBack(val)
}

func (q *Queue) Pop() interface{} {
	front := q.v.Front()
	if front != nil {
		return q.v.Remove(front)
	}
	return nil
}

func NewQueue() *Queue {
	return &Queue{list.New()}
}

func main() {
	// list를 이용해 큐 자료구조 구현하기
	queue := NewQueue()

	for i := 1; i < 5; i++ {
		queue.Push(i)
	}

	v := queue.Pop()
	for v != nil {
		fmt.Printf("%v ->", v)
		v = queue.Pop()
	}
}
