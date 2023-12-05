package main

import "fmt"

func main() {
	// 기본적인 배열 선언 방법
	var t [5]float64 = [5]float64{24.0, 25.9, 27.8, 26.9, 26.2}

	for i := 0; i < 5; i++ {
		fmt.Println(t[i])
	}

	// 다양한 배열 선언 방법 예시
	// 값을 지정하지 않으면 기본값이 자동으로(0, 0.0...)
	// var nums [5]int
	// 선언 대입문(:=) 가능
	// days := [3]string{"monday", "tuesday", "wednesday"}
	// 앞의 두 값만 지정되고 나머지는 기본값
	// var temps [5]float64 = [5]float64{24.3, 26.7}
	// 인덱스에 해당하는 값만 지정됨
	// var s = [5]int{1: 10, 3: 30}
	// [...] 으로 선언 시 뒤의 배열의 크기만큼 고정
	// x := [...]int{10, 20, 30}
}
