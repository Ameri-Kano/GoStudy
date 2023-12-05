package main

import "fmt"

func main() {
	// for 초기문; 조건문; 후처리문
	// 초기문과 후처리문 모두 생략 가능하지만 ;을 표시해줘야 함
	// 조건문만 있으면 ;도 생략 가능
	// go는 다른 언어의 while 문의 역할도 for 문이 수행함
	for i := 0; i < 10; i++ {
		fmt.Print(i, ", ")
	}
}
