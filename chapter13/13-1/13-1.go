package main

import "fmt"

// House 라는 이름의 구조체 선언
//
//	type 이름 struct {
//			필드명 타입
//			...
//			필드명 타입
//	}
type House struct {
	Address  string
	Size     int
	Price    float64
	Category string
}

func main() {
	var house House
	house.Address = "서울시 은평구 ..."
	house.Size = 32
	house.Price = 10
	house.Category = "아파트"

	// 단순하게 출력 시 중괄호({ }) 안에 필드를 순서대로 나열 (%v 서식의 결과와 같음)
	fmt.Println(house)

	// 대입했던 것처럼 값에 접근하여 출력 가능
	fmt.Printf("주소:%s 사이즈:%d평 가격:%v억 종류:%s\n",
		house.Address, house.Size, house.Price, house.Category)
}
