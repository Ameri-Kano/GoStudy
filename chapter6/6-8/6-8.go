package main

import (
	"fmt"
	"math"
)

func equal(a, b float64) bool {
	return math.Nextafter(a, b) == b
}

func main() {
	var a float64 = 0.1
	var b float64 = 0.2
	var c float64 = 0.3

	// math.Nextafter() : a에서 b를 향해 1비트만 간다
	// c -> a+b 를 향해 1비트만 증가시키면 결과가 같아짐
	// 이런 방법으로 아주 미세한 오차를 무시 가능
	// 0.299999999999999989 == 0.300000000000000044 : true
	fmt.Printf("%0.18f == %0.18f : %v\n", c, a+b, equal(a+b, c))
}
