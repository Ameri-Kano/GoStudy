package main

import "fmt"

func main() {
	var a = 123
	var b = 456
	var c = 123456789

	// 0이 앞에 들어가면 빈자리는 0으로 채워짐
	// -가 붙으면 왼쪽부터 (기본 오른쪽)
	fmt.Printf("%5d, %5d\n", a, b)
	fmt.Printf("%05d, %05d\n", a, b)
	fmt.Printf("%-5d, %-05d\n", a, b)
	// 서식보다 변수의 원래 자리수가 많다고 생략되지는 않음
	fmt.Printf("%5d, %5d\n", c, c)
	fmt.Printf("%05d, %05d\n", c, c)
	fmt.Printf("%-5d, %-05d\n", c, c)
}
