package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("test.txt")

	if err != nil {
		fmt.Println("Failed to create a file", err)
		return
	}

	// defer 사용시 바로 실행되지 않고(지연) 함수가 종료되기 직전에 실행됨
	defer fmt.Println("반드시 호출됩니다")
	defer f.Close()
	defer fmt.Println("파일을 닫았습니다")

	fmt.Println("파일에 Hello World를 씁니다")
	fmt.Fprintln(f, "Hello World")
}
