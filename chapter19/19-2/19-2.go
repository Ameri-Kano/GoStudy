package main

import "fmt"

type myInt int

// Add 메서드
// 메서드는 패키지 내부의 타입만 리시버로 받을 수 있다
// int 를 myInt 라는 별칭 타입으로 지정해 메서드 구현
func (a myInt) Add(b int) myInt {
	rst := int(a) + b
	return myInt(rst)
}

func addFunc(m int, a int) myInt {
	rst := m + a
	return myInt(rst)
}

func main() {
	var a myInt
	var b int
	a = a.Add(10)
	b = int(addFunc(b, 10))

	fmt.Println(a)
	fmt.Println(b)
}
