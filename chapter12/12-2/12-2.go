package main

const Y int = 3

func main() {
	x := 5
	// 변수의 경우 배열의 길이 값으로 정할 수 없음(에러 발생)
	a := [x]int{1, 2, 3, 4, 5}
	// Y가 상수(const) 이므로 길이 값 지정 가능
	b := [Y]int{1, 2, 3}

	// 상수(10) 이기 때문에 에러 발생하지 않음
	var c [10]int
}
