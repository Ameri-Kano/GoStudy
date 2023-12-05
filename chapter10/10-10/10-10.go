package main

import "fmt"

func main() {
	a := 3

	switch a {
	case 1:
		fmt.Println("a == 1")
		// break 를 따로 작성할 필요 없이 Go 에선 하나의 case 가 실행되면 switch 끝
		break
	case 2:
		fmt.Println("a == 2")
	case 3:
		fmt.Println("a == 3")
		// 그 다음 case까지 실행하고 싶다면 fallthrough 사용
		// 그러나 사용을 권장하지는 않음
		fallthrough
	case 4:
		fmt.Println("a == 4")
	default:
		fmt.Println("a != 1, 2, 3, 4")
	}
}
