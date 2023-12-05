package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	b := [5]int{500, 400, 300, 200, 100}

	for i, v := range a {
		fmt.Printf("a[%d] = %d\n", i, v)
	}

	fmt.Println()
	for i, v := range b {
		fmt.Printf("b[%d] = %d\n", i, v)
	}

	// Go 에서 대입 연산자는 우변의 값을 좌변의 메모리 공간에 복사한다
	// 복사되는 크기는 타입의 크기와 같으며 배열의 대입도 마찬가지
	// b = a를 하면 a의 값을 b의 메모리 공간에 복사한다.
	// 그러므로 배열의 대입은 전체 배열의 복사로 동작한다.
	// !!!중요!!! 복사하는 배열과 값이 대입되는 배열의 크기와 타입이 꼭 같아야 한다!
	b = a
	fmt.Println()
	for i, v := range b {
		fmt.Printf("b[%d] = %d\n", i, v)
	}
}
