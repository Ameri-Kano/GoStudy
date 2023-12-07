package main

import "fmt"

func f() {
	fmt.Println("f() 함수 시작")

	// recover() 는 발생한 패닉을 반환하고 복구함
	// 패닉이 복구되면 프로그램이 종료되지 않음
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic 복구 - ", r)
		}
	}()

	g()
	fmt.Println("f() 함수 끝")
}

func g() {
	fmt.Printf("9/3 = %d\n", h(9, 3))
	fmt.Printf("9/0 = %d\n", h(9, 0))
}

func h(a, b int) int {
	if b == 0 {
		panic("제수는 0일 수 없습니다.")
	}
	return a / b
}

func main() {
	f()
	// 복구되지 않으면 아래 코드가 실행되지 않음
	fmt.Println("프로그램이 계속 실행됨")
}
