package main

import (
	"fmt"
	"os"
)

type Writer func(string)

type WriterInterface interface {
	Write(string)
}

func writeHello(writer Writer) {
	writer("Hello World")
}

// 외부에서 로직을 주입하는 것을 의존성 주입(DI) 라고 한다
// 함수 또는 인터페이스로 구현 가능
func writerHello2(writer WriterInterface) {
	writer.Write("Hello World")
}

func main() {
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Failed to create a file")
		return
	}
	defer f.Close()

	writeHello(func(msg string) {
		fmt.Fprintln(f, msg)
	})
}
