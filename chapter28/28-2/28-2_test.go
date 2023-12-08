package main

import "testing"

// BenchmarkXxxx(b *testing.B) 의 형태로 작성
// 결과는 벤치마크 이름 / 반복 횟수 / 실행 시간 순서로 나옴
// 벤치마크를 통해 성능 측정이 쉽게 가능
func BenchmarkFibonacci1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacci1(20)
	}
}

func BenchmarkFibonacci2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacci2(20)
	}
}
