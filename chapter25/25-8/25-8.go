package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	// 컨텍스트 기본 활용
	// 취소 가능한 컨텍스트 생성 (컨텍스트, 취소함수)
	// 기본 컨텍스트는 context.BackGround()
	wg.Add(1)
	ctx, cancel := context.WithCancel(context.Background())

	// 5초 후에 작업 취소를 알리면
	// ctx.Done() 채널에 신호를 보냄
	// 신호를 받으면 프로그램 종료
	go PrintEverySecond(ctx)
	time.Sleep(5 * time.Second)
	cancel()

	wg.Wait()
}

func PrintEverySecond(ctx context.Context) {
	tick := time.Tick(time.Second)
	for {
		select {
		case <-ctx.Done():
			wg.Done()
			return
		case <-tick:
			fmt.Println("Tick")
		}
	}
}
