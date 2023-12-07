package main

import "fmt"

func main() {
	m := make(map[int]int)
	m[1] = 0
	m[2] = 2
	m[3] = 3

	// delete(맵, 키) 를 이용해 키에 해당하는 값 삭제
	delete(m, 3)
	delete(m, 4)

	// 키에 해당하는 값이 존재하면 값, true 반환
	// 존재하지 않으면 기본값, false 반환
	v, ok := m[3]
	fmt.Println(v, ok)
	v, ok = m[1]
	fmt.Println(v, ok)
	// fmt.Println(m[1])
}
