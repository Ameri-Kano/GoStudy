package main

import "fmt"

func CaptureLoop() {
	f := make([]func(), 3)
	fmt.Println("ValueLoop")
	// 변수가 캡쳐된 상태이므로 출력하면 3 3 3 이 반복적으로 나옴
	i := 0
	for ; i < 3; i++ {
		f[i] = func() {
			fmt.Println(i)
		}
	}

	// i = 3

	for j := 0; j < 3; j++ {
		f[j]()
	}
}

func CaptureLoop2() {
	f := make([]func(), 3)
	fmt.Println("ValueLoop2")

	for i := 0; i < 3; i++ {
		v := i
		f[i] = func() {
			fmt.Println(v)
		}
	}

	for i := 0; i < 3; i++ {
		f[i]()
	}
}

func main() {
	CaptureLoop()
	CaptureLoop2()
}
