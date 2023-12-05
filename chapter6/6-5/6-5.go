package main

import "fmt"

func main() {
	var x int8 = 127

	// 127 < 127 + 1 일텐데 왜 false 가 나오는가?
	// 1바이트 정수값 중 최댓값(01111111)에 1을 더하면 10000000 = -128
	// 범위를 벗어나 오버플로우 발생
	fmt.Printf("%d < %d+1: %v\n", x, x, x < x+1)
	fmt.Printf("x\t= %4d, %08b\n", x, x)
	fmt.Printf("x+1\t= %4d, %08b\n", x+1, uint8(x+1))
}
