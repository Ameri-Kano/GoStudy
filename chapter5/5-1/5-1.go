package main

import "fmt"

func main() {
	var a int = 10
	var b int = 20

	var f float64 = 32799438743.8297

	// Print() - 줄바꿈 없이 (입력값 사이에 빈칸 X)
	// Println() - 출력 후 줄바꿈 (입력값 사이에 빈칸 O)
	// Printf() - 서식에 맞춰 출력
	fmt.Print("a:", a, "b:", b)
	fmt.Println("a:", a, "b:", b, "f:", f)
	fmt.Printf("a: %d b: %d f: %f\n", a, b, f)
}
