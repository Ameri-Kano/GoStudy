package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	ch := make(chan int)

	wg.Add(1)
	go square(&wg, ch)

	for i := 0; i < 10; i++ {
		ch <- i * 2
	}
	wg.Wait()
}

func square(wg *sync.WaitGroup, ch chan int) {
	// Tick() 은 정한 매 시간 간격 으로 신호를 전달하는 채널 반환
	tick := time.Tick(time.Second)
	// After() 는 지금부터 일정 시간 이후에 신호를 전달하는 채널 반환
	terminate := time.After(10 * time.Second)

	// select 를 이용해 여러 채널을 동시에 기다릴 수 있음
	// for 무한루프를 이용해야 종료되기 전까지 반복해서 대기함
	for {
		select {
		case <-tick:
			fmt.Println("Tick")
		case <-terminate:
			fmt.Println("Terminated")
			wg.Done()
			return
		case n := <-ch:
			fmt.Println("Square:", n*n)
			time.Sleep(time.Second)
		}
	}
}
