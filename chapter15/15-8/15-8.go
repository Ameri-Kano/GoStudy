package main

import "fmt"

func main() {
	str := "Hello 월드"
	// 문자열은 고정 배열로 변환할 수 없다
	// [] 로 선언시 슬라이스로 변환 가능(슬라이스에 대해서는 18장)
	// rune은 int32의 별칭
	arr := []rune(str)

	// 같은 문자열을 깨지지 않고 정상적으로 출력
	for i := 0; i < len(arr); i++ {
		fmt.Printf("타입:%T 값:%d 문자값:%c\n", arr[i], arr[i], arr[i])
	}
}
