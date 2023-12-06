package main

import "fmt"

func main() {
	var str string = "Hello World"
	var slice []byte = []byte(str)

	// 문자열은 한번 생성되면 변하지 않는다
	// 슬라이스의 경우 문자열을 바탕으로 새로운 메모리 공간에 생성한다
	slice[2] = 'a'
	fmt.Println(str)
	fmt.Printf("%s\n", slice)
}
