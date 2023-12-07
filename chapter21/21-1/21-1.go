package main

import "fmt"

func sum(nums ...int) int {
	sum := 0

	fmt.Printf("nums 타입: %T\n", nums)
	for _, v := range nums {
		sum += v
	}
	return sum
}

func main() {
	// ...Type 을 사용하여 해당 Type의 값을 제한 없이 받을 수 있음
	// 넘어갈 때 타입은 []Type (슬라이스) 가 됨
	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(sum(10, 20))
	fmt.Println(sum())
}
