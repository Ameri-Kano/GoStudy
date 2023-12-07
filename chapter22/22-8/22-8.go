package main

import "fmt"

const M = 10

func hash(d int) int {
	return d % M
}

func main() {
	m := [M]string{}

	// 맵의 원리
	// 해시 함수를 이용하여 저장
	// 그러나 만약 해시값이 같으면 원래 있던 값이 덮어씌워 질수도 있음
	// 이를 해결하려면 각 인덱스 위치에 값이 아닌 리스트를 저장하면 됨
	m[hash(23)] = "송하나"
	m[hash(259)] = "백두산"

	fmt.Printf("%d = %s\n", 23, m[hash(23)])
	fmt.Printf("%d = %s\n", 259, m[hash(259)])
}
