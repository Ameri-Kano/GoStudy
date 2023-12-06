package main

import "fmt"

func main() {
	str := "Hello 월드"

	// 이렇게 출력하면 한글 문자는 깨지게 됨
	// 영어는 1바이트 표현이나 한글은 3바이트로 표현하기 떄문
	// len(문자열) 의 경우 문자열의 글자 길이가 아니라 차지하는 바이트 길이를 반환
	for i := 0; i < len(str); i++ {
		fmt.Printf("타입:%T 값:%d 문자값:%c\n", str[i], str[i], str[i])
	}
}
