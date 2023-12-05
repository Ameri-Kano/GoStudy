package main

import "fmt"

func main() {
	// 기본적인 선언 형태
	var a int = 3
	// 값이 없으면 기본값 (int의 경우 0)
	var b int
	// 타입을 생략하면 우변의 값 타입
	var c = 4
	// var 생략하여 표현 가능
	d := 5

	fmt.Println(a, b, c, d)
}
