package main

import "fmt"

func main() {
	// 2차원 배열 선언
	a := [2][5]int{
		{1, 2, 3, 4, 5},
		{5, 6, 7, 8, 9},
	}

	// 배열이므로 range 를 이용하여 순회 가능
	for _, arr := range a {
		for _, v := range arr {
			fmt.Print(v, " ")
		}
		fmt.Println()
	}
}
