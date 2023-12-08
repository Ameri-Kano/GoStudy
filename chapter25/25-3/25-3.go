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
	// 해결하기 위해서 close(ch) 를 통해 채널을 닫아주기
	close(ch)
	wg.Wait()
}

func square(wg *sync.WaitGroup, ch chan int) {
	// range 를 이용하여 채널의 데이터를 무한히 대기하며 받을 수 있음
	// 채널을 따로 닫아주지 않으면 무한히 대기하므로 데드락 발생 가능
	for n := range ch {
		fmt.Println("Square:", n*n)
		time.Sleep(time.Second)
	}
	wg.Done()
}
