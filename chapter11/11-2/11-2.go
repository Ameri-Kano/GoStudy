package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1
	// for true 면 무한루프 수행
	// true 를 생략해도 무한루프 가능
	for {
		time.Sleep(time.Second)
		fmt.Println(i)
		i++
	}
}
