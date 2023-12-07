package main

import (
	"container/ring"
	"fmt"
)

func main() {
	// 링 자료구조(원형 리스트)
	// 일정한 갯수만 사용하고 오래된 요소가 지워져도 되는 경우에 사용
	// 1. 문서 편집기 등에서 일정 개수의 명령을 저장하고 실행 취소
	// 2. 고정 크기 버퍼
	// 3. 게임 등에서 최근 플레이 10초를 다시 리플레이 하는 것 처럼 고정된 길이의 리플레이 기능
	r := ring.New(5)

	n := r.Len()
	// A B C D E 저장
	for i := 0; i < n; i++ {
		r.Value = 'A' + i
		r = r.Next()
	}
	// A부터 시작하여 원래 방향대로
	for j := 0; j < n; j++ {
		fmt.Printf("%c ", r.Value)
		r = r.Next()
	}
	// A부터 시작하여 반대로 돌기
	fmt.Println()
	for j := 0; j < n; j++ {
		fmt.Printf("%c ", r.Value)
		r = r.Prev()
	}
}
