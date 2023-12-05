package main

import "fmt"

func main() {
	a := 3

	// switch - case
	// case 안에는 괄호로 감싸지 않아도 여러 줄 작성 가능
	switch a {
	case 1:
		fmt.Println("a == 1")
	case 2:
		fmt.Println("a == 2")
	case 3:
		fmt.Println("a == 3")
	default:
		fmt.Println("a != 1, 2, 3", a)
	}
}
