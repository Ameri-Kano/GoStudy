package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)

	var a int
	var b int

	n, err := fmt.Scanln(&a, &b)
	if err != nil {
		fmt.Println(err)
		// 표준 입력에서 특정 문자가 나올때까지 읽기
		// 개행문자가 입력될때까지 읽어서 표준 입력 스트림을 비울 수 있음
		// 비우지 않으면 에러 발생 후 다음 입력에서 의도치 않은 입력 발생 가능
		stdin.ReadString('\n')
	} else {
		fmt.Println(n, a, b)
	}

	n, err = fmt.Scanln(&a, &b)
	if err != nil {
		fmt.Println(err)
		stdin.ReadString('\n')
	} else {
		fmt.Println(n, a, b)
	}
}
