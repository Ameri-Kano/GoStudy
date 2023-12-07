package main

import "fmt"

func divide(a, b int) {
	if b == 0 {
		// 패닉이 발생하면 바로 강제 종료
		// 어디에서 패닉이 발생했는지 알려주므로 빠르게 문제 발견 가능
		panic("b는 0일 수 없습니다")
	}
	fmt.Printf("%d / %d = %d\n", a, b, a/b)
}

func main() {
	divide(9, 3)
	divide(9, 0)
}
