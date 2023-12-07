package main

import (
	"fmt"
	"time"
)

func PrintHangul() {
	hanguls := []rune{'가', '나', '다', '라', '마', '바', '사'}
	for _, v := range hanguls {
		time.Sleep(300 * time.Millisecond)
		fmt.Printf("%c ", v)
	}
}

func PrintNumbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

func main() {
	// go 키워드를 사용해 각각의 새로운 고루틴에서 함수 실행
	go PrintHangul()
	go PrintNumbers()

	// 만약 main에서 3초를 기다리지 않았다면
	// 프로그램은 main에서 시작하고 끝나므로 상관 없이 종료됨
	time.Sleep(3 * time.Second)
}
