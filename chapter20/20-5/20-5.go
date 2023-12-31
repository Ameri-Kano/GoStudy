package main

import "fmt"

func PrintVal(v interface{}) {
	switch t := v.(type) {
	case int:
		fmt.Printf("v is int %d\n", int(t))
	case float64:
		fmt.Printf("v is float64 %f\n", float64(t))
	case string:
		fmt.Printf("v is string %s\n", string(t))
	default:
		fmt.Printf("Not supported type %T:%v\n", t, t)
	}
}

type Student struct {
	Age int
}

func main() {
	// interface{} 라고 작성하면 빈 인터페이스가 된다
	// 함수에서 빈 인터페이스를 인수로 받으면 모든 타입을 쓸 수 있음
	PrintVal(10)
	PrintVal(3.14)
	PrintVal("Hello")
	PrintVal(Student{15})
}
