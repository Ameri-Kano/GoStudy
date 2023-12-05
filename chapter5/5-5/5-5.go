package main

import "fmt"

func main() {
	var a int
	var b int

	// 표준 입력 (변수의 주소를 전달해야함)
	// 주소를 전달해야 입력받은 값을 주소에 넣어서 저장 가능
	n, err := fmt.Scanln(&a, &b)

	// 에러가 없으면 err 은 nil 이 됨
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(n, a, b)
	}
}
