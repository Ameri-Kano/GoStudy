package main

import "fmt"

// printNo 는 재귀함수
// 0이 되면 재귀 탈출(안하면 무한 호출)
// 스택과 같이 가장 나중에 호출된 함수부터 실행
func printNo(n int) {
	if n == 0 {
		return
	}
	fmt.Println(n)
	printNo(n - 1)
	fmt.Println("After", n)
}

func main() {
	// 3
	// 2
	// 1
	// After 1
	// After 2
	// After 3
	printNo(3)
}
