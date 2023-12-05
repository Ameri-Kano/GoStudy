package main

import "fmt"

// Divide 함수는 7-4 와 같음
// 반환 값에 변수 이름을 지정하면 함수 범위 내에서 사용 가능
// 이런 경우 따로 return 에 반환값을 명시할 필요 없음
func Divide(a, b int) (result int, success bool) {
	if b == 0 {
		result = 0
		success = false
		return
	}
	result = a / b
	success = true
	return
}

func main() {
	c, success := Divide(9, 3)
	fmt.Println(c, success)
	// 0, false
	d, success := Divide(9, 0)
	fmt.Println(d, success)
}
