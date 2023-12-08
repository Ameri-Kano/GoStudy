package main

import (
	"context"
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)

	// 컨텍스트에 값 추가
	ctx := context.WithValue(context.Background(), "number", 9)

	go square(ctx)

	wg.Wait()
}

func square(ctx context.Context) {
	// 컨텍스트에서 값 읽기
	if v := ctx.Value("number"); v != nil {
		n := v.(int)
		fmt.Printf("Square:%d", n*n)
	}
	wg.Done()
}
