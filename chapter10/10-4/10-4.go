package main

import "fmt"

func main() {

	day := "thursday"

	// 다음과 같이 , 로 한 줄의 case에 여러 값을 넣을 수 있음
	switch day {
	case "monday", "tuesday":
		fmt.Println("월, 화요일은 수업가는 날")
	case "wednesday", "thursday", "friday":
		fmt.Println("수, 목, 금요일은 실습가는 날")
	}
}
