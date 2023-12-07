package main

import "fmt"

func main() {
	// 맵(map) 자료구조의 기본 사용법
	// make(map[키타입]값타입)
	m := make(map[string]string)

	m["이화랑"] = "서울시 광진구"
	m["송하나"] = "서울시 강남구"
	m["백두산"] = "부산시 사하구"
	m["최번개"] = "전주시 덕진구"
	// 값 변경 가능
	m["최번개"] = "청주시 상당구"

	fmt.Printf("송하나의 주소는 %s입니다.\n", m["송하나"])
	fmt.Printf("백두산의 주소는 %s입니다.\n", m["백두산"])
}
