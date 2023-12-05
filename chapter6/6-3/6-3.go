package main

import "fmt"

func main() {
	var x int8 = 4
	var y int8 = 64

	// 시프트 연산자로 한 칸 옮긴다고 무조건 기존 수의 2배가 되지 않는다.
	// y의 경우 비트를 두 칸 밀면 1이 전부 사라져 0이 되기 때문
	fmt.Printf("x:%08b x<<2:%08b x<<2: %d\n", x, x<<2, x<<2)
	fmt.Printf("y:%08b y<<2:%08b y<<2: %d\n", y, y<<2, y<<2)
}
