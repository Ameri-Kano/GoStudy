package main

import (
	"container/list"
	"fmt"
)

type Stack struct {
	v *list.List
}

func NewStack() *Stack {
	return &Stack{list.New()}
}

func (s *Stack) Push(val interface{}) {
	s.v.PushBack(val)
}

func (s *Stack) Pop() interface{} {
	back := s.v.Back()
	if back != nil {
		return s.v.Remove(back)
	}
	return nil
}

func main() {
	// list를 이용해 스택 자료구조 구현하기
	// 흔히 스택은 list보단 배열로 구현한다
	stack := NewStack()
	books := [5]string{"어린왕자", "겨울왕국", "노인과바다", "빨간머리앤", "짱구"}

	for i := 0; i < 5; i++ {
		stack.Push(books[i])
	}

	val := stack.Pop()
	for val != nil {
		fmt.Printf("%v ->", val)
		val = stack.Pop()
	}
}
