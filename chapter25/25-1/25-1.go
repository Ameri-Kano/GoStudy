package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// 채널을 이용하여 동시성 프로그래밍 구현 가능
	var wg sync.WaitGroup
	// 채널 생성 방법
	// make(chan 타입)
	ch := make(chan int)

	wg.Add(1)
	go square(&wg, ch)
	// 채널에 데이터 넣는 방법
	// 채널 <- 데이터
	ch <- 9
	wg.Wait()
}

func square(wg *sync.WaitGroup, ch chan int) {
	// 채널에서 데이터 가져오기
	// 변수 <- 데이터
	n := <-ch

	time.Sleep(time.Second)
	fmt.Println("Square:", n*n)
	wg.Done()
}
