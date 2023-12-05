package main

import "fmt"

func main() {
	temp := 33

	// if - elseif - else 문 기본 구조
	if temp > 28 {
		fmt.Println("에어컨을 킨다")
	} else if temp <= 3 {
		fmt.Println("히터를 킨다")
	} else if temp <= 18 {
		fmt.Println("나가자!")
	} else {
		fmt.Println("덥다")
	}
}
