package main

import "fmt"

func main() {
	var t [5]float64 = [5]float64{24.0, 25.9, 27.8, 26.9, 26.2}

	// range를 이용하여 배열 요소 전체 순회(Java 의 for-each처럼)
	// for 인덱스값 요소값 := range 배열
	for i, v := range t {
		fmt.Println(i, v)
	}

	// 만약 요소만 필요하면 인덱스 변수는 _를 이용해 무효화할 수 있음
	// Go 에서는 변수를 선언했으면 사용해야 하기 때문(안할 시 에러)
	for _, v := range t {
		fmt.Println(v)
	}
}
