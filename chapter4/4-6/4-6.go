package main

import "fmt"

// 패키지 전역변수
var g int = 10

func main() {
	var m int = 20

	// 변수의 범위(그 변수가 속한 중괄호)를 벗어나면 그 변수는 사라짐
	// 아래 코드에서 s는 사라진 상태이므로 에러 발생
	{
		var s int = 50
		fmt.Println(m, s, g)
	}

	// m = s + 20
}
