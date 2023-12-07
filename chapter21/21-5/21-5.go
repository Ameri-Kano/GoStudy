package main

import "fmt"

func main() {
	i := 0

	// 함수 리터럴의 경우 함수 외부의 변수를 사용할 수 있음
	// 외부 변수 캡쳐는 값 복사가 아닌 포인터 복사
	f := func() {
		i += 10
	}

	i++
	f()

	fmt.Println(i)
}
