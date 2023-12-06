package main

import "fmt"

func main() {
	// 포인터 : 변수가 저장되는 메모리 주소값을 가지는 변수
	var a int = 500
	var p *int

	p = &a

	fmt.Printf("p의 값: %p\n", p)
	fmt.Printf("p가 가리키는 메모리 공간의 값: %d\n", *p)

	*p = 100
	fmt.Printf("a의 값: %d\n", a)
}
