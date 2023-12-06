package main

import "fmt"

func main() {
	var a int = 10
	var b int = 20

	var p1 *int = &a
	var p2 *int = &a
	var p3 *int = &b

	// p1과 p2는 같은 변수의 주소를 가지므로 값이 같다
	// p2와 p3는 다른 변수의 주소를 가지므로 값이 다르다
	fmt.Printf("p1 == p2 %v\n", p1 == p2)
	fmt.Printf("p2 == p3 %v\n", p2 == p3)
}
