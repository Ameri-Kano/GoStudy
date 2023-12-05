package main

import "fmt"

func main() {
	temp := 18

	// case에 조건문이 들어갈 수도 있음
	// switch true의 경우 case가 true가 되는 문장 실행
	switch true {
	case temp < 10, temp > 30:
		fmt.Println("바깥 활동하기 좋은 날씨가 아닙니다")
	case temp > 10 && temp < 20:
		fmt.Println("약간 추울 수 있으니 가벼운 겉옷을 준비하세요")
	// 조건을 만족하지만 이미 위 로직이 실행되어 switch문이 종료됨
	// 따라서 아래 조건을 검사하지 않음
	case temp > 15 && temp < 25:
		fmt.Println("야외 활동하기 좋은 날씨입니다")
	default:
		fmt.Println("따듯합니다")
	}
}
