package main

import "fmt"

func main() {
	str1 := "Hello"
	str2 := "World"

	// 문자열을 서로 더하면 이어붙이기 가능
	str3 := str1 + " " + str2
	fmt.Println(str3)
	// 복합 대입 연산자로도 가능
	str1 += " " + str2
	fmt.Println(str1)
}
