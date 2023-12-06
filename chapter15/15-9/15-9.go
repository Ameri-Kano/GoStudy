package main

import "fmt"

func main() {
	str := "Hello 월드"

	// range 를 이용하여 앞의 예제들과 같은 결과 출력 가능
	// 한 글자씩 순회
	for _, v := range str {
		fmt.Printf("타입:%T 값:%d 문자값:%c\n", v, v, v)
	}
}
