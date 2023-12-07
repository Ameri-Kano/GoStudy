package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func mul(a, b int) int {
	return a * b
}

// OpFn 함수 타입도 별칭으로 지정 가능
type OpFn func(int, int) int

func getOperator(op string) OpFn {
	if op == "+" {
		return add
	} else if op == "*" {
		return mul
	} else {
		return nil
	}
}

func main() {
	// 함수를 변수로 가질 수 있다.
	// 함수의 시작 주소를 변수가 가지게 됨
	// operator = add()
	var operator OpFn
	operator = getOperator("")

	var result = operator(3, 4)
	fmt.Println(result)
}
