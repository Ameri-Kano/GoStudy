package main

import "fmt"

func main() {
	a := 1
	b := 1

	// label(레이블)을 이용하여 더 간단하게 반복문 탈출 가능
	// 그러나 혼동을 줄 수 있고 실행 중 메모리에서 버그가 발생할 수도 있음
	// 그러므로 가급적 사용하지 않는 것을 권장(앞의 예제에서처럼 flag 변수 사용하기)
OuterFor:
	for ; a <= 9; a++ {
		for b = 1; b <= 9; b++ {
			if a*b == 45 {
				break OuterFor
			}
		}
	}
	fmt.Printf("%d * %d = %d\n", a, b, a*b)
}
