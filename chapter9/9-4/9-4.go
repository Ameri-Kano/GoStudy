package main

import "fmt"

var cnt int = 0

func IncreaseAndReturn() int {
	fmt.Println("IncreaseAndReturn()", cnt)
	cnt++
	return cnt
}

func main() {
	// 조건에 의해서라면 출력되어야 할 것
	// && 의 경우 왼쪽이 false면 오른쪽 상관없이 false
	// 따라서 오른쪽을 체크하지 않으면서 함수가 실행되지도 않음
	// || 의 경우 왼쪽이 true면 오른쪽 상관없이 true
	// 마찬가지로 오른쪽을 체크하지 않음
	// 이러한 쇼트 서킷이 존재
	if false && IncreaseAndReturn() < 5 {
		fmt.Println("1 증가")
	}
}
