package main

import "fmt"

func main() {
	var a int16 = 3456 // int16 - 2byte 정수
	var b int8 = int8(a)

	// 왜 b는 -128이 나올까?
	// 16비트 정수가 8비트 정수에 담기면서 앞의 8비트가 없어짐
	// 00001101 10000000 -> 10000000 = -128
	fmt.Println(a, b)
}
