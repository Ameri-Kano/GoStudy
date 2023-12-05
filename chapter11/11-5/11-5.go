package main

import "fmt"

func main() {
	// 이중 반복문 활용하여 다양한 도형 그리기
	// 1.
	// *
	// **
	// ***
	// ****
	// *****
	for i := 0; i < 5; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	fmt.Println()
	// 2.
	// *****
	// ****
	// ***
	// **
	// *
	for i := 5; i > 0; i-- {
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	fmt.Println()
	// 3.
	//   *
	//  ***
	// *****
	//*******
	for i := 0; i < 4; i++ {
		for j := 0; j < 3-i; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < 2*i+1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	fmt.Println()
	// 4.
	//   *
	//  ***
	// *****
	//  ***
	//   *
	for i := 0; i < 3; i++ {
		for j := 0; j < 2-i; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < 2*i+1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	for i := 1; i >= 0; i-- {
		for j := 0; j < 2-i; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < 2*i+1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	fmt.Println()
}
