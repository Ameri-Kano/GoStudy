package main

import "fmt"

func main() {
	a := 3 // int -> 64bit = int64
	var b float64 = 3.5

	// 무조건 같은 타입만 연산이 가능하다
	// 같은 정수형이어도 int16 과 int64는 서로 연산 불가능
	var c int = int(b)
	d := float64(a) * b

	var e int64 = 7
	f := a * int(e)

	fmt.Println(a, b, c, d, e, f)
}
