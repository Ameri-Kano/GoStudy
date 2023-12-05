package main

import "fmt"

func main() {
	var a float32 = 1234.523
	var b float32 = 3456.123

	// float32 는 7자리까지만 표시가능
	// 숫자가 제한되어 오차가 커지는 문제 발생
	// c 4.266663e+06
	// d 1.2799989e+07
	var c float32 = a * b
	var d float32 = c * 3

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
