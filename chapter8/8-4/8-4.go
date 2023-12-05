package main

import "fmt"

const PI = 3.14
const FloatPI float64 = 3.14

func main() {
	// 상수에 타입을 지정하지 않으면 단순한 숫자로 동작
	// 타입을 지정할 시 일반적인 타입 검사 수행
	var a int = PI * 100
	// var b int = FloatPI * 100

	fmt.Println(a)
	// fmt.Println(b)
}
