package main

import "fmt"

func main() {
	var a float64 = 0.1
	var b float64 = 0.2
	var c float64 = 0.3

	// 0.1 + 0.2 = 0.3 일텐데 왜 false인가?
	// 컴퓨터는 소수를 부호, 지수, 소수부 비트로 나누어 표현
	// 그러므로 정확하게 0.300000... 이 될 수 없음(그러나 0.3에는 매우 가까움)
	fmt.Printf("%f + %f == %f : %v\n", a, b, c, a+b == c)
}
