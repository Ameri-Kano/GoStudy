package main

import "fmt"

func main() {
	// 중첩 for문
	// 가로 5, 세로 3 사각형 그리기
	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
