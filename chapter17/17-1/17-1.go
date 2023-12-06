package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// time 패키지는 시간과 관련된 기능 제공
	// time.Now() : 현재 시각 제공
	// UnixNano() : 1970년 1월 1일부터 현재까지 경과된 시간을 nanosecond 단위로 반환
	// t := time.Now()
	// 그러나 현재 버전에서는 Seed 설정이 따로 필요없는 듯 함(deprecated 되었음)
	// rand.Seed(t.UnixNano())

	// rand.Intn(n) : 0부터 n-1 까지의 무작위 정수 반환
	for i := 0; i < 10; i++ {
		fmt.Println(rand.Intn(100))
	}
}
