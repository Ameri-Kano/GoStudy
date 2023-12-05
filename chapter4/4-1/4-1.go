package main

import "fmt"

func main() {
	// 변수 이름 자료형
	var a int = 10
	var msg string = "Hello Variable"

	// 값 변경
	a = 20
	msg = "Good Morning"

	// 한 칸 띄우고 출력 "20 Good Morning"
	fmt.Println(msg, a)
}
