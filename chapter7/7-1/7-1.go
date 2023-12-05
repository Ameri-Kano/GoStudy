package main

import "fmt"

// Add 함수는 두 정수를 더한 값을 반환
// func 함수이름(파라미터) 반환타입
func Add(a int, b int) int {
	return a + b
}

func main() {
	c := Add(3, 6)
	fmt.Println(c)
}
