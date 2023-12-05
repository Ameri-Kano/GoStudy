package main

import "fmt"

// Divide 함수는 a를 b로 나눈 수를 반환함
// 한 번에 여러 값을 반환 가능함 (괄호로 묶어주고 타입)
// 여러 파라미터가 같은 값일 경우 한 번만 타입 명시 가능
func Divide(a, b int) (int, bool) {
	if b == 0 {
		return 0, false
	}
	return a / b, true
}

func main() {
	// 호출이 끝나면 3, true 로 바뀜
	// 값의 대입도 마찬가지로 여러 값을 한번에 가능
	c, success := Divide(9, 3)
	fmt.Println(c, success)
	// 0, false
	d, success := Divide(9, 0)
	fmt.Println(d, success)
}
