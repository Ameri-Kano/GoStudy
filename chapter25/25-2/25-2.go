package main

import (
	"fmt"
	"time"
)

func main() {
	// 채널 생성시 기본적으로 크기가 0인 채널이 만들어짐
	// 데이터를 빼 가는 코드가 없으므로 모든 고루틴이 영원히 대기(데드락)
	// 데이터 보관 영역을 만드려면 make(chan 타입, n) 으로 선언
	ch := make(chan int, 2)

	// go square()
	ch <- 9
	fmt.Println("Never print")
}

func square() {
	for {
		time.Sleep(2 * time.Second)
		fmt.Println("sleep")
	}
}
