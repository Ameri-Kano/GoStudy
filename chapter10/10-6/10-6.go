package main

import "fmt"

func getMyAge() int {
	return 25
}

func main() {
	// if와 마찬가지로 switch 도 초기문 사용 가능
	switch age := getMyAge(); age {
	case 10:
		fmt.Println("Teenage")
	case 33:
		fmt.Println("Pair 3")
	default:
		fmt.Println("My age is", age)
	}

	// age 는 switch 문 에서 선언되었으므로 if와 마찬가지로 벗어나면 사라짐
	// 아래 코드는 같은 이유로 에러 발생
	// fmt.Println("age is", age)
}
