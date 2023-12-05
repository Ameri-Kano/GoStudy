package main

import "fmt"

func main() {
	a := 1
	b := 1
	found := false
	// 곱한 결과가 45일때 탈출
	// 이중 반복문이므로 flag 변수를 확인해 한번 더 탈출
	for ; a <= 9; a++ {
		for b = 1; b <= 9; b++ {
			if a*b == 45 {
				found = true
				break
			}
		}
		if found {
			break
		}
	}
	fmt.Printf("%d * %d = %d\n", a, b, a*b)
}
