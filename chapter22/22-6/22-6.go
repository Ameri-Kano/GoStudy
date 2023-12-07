package main

import "fmt"

type Product struct {
	Name  string
	Price int
}

func main() {
	m := make(map[int]Product)

	m[16] = Product{"볼펜", 500}
	m[46] = Product{"지우개", 200}
	m[78] = Product{"자", 1000}
	m[345] = Product{"샤프", 3000}
	m[897] = Product{"샤프심", 500}
	// 맵 역시 range 를 사용해 순회 가능 (키, 값)
	// 그러나 자료의 순서를 보장하지 않음(Go 는 해시맵 방식 사용)
	for k, v := range m {
		fmt.Println(k, v)
	}
}
