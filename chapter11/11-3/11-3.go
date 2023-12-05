package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("입력")
		var number int
		// 할당되는 변수를 무조건 사용해야 한다
		// 할당시키지 않으려면 _ 사용
		_, err := fmt.Scanln(&number)
		if err != nil {
			fmt.Println("숫자를 입력하세요")

			// 키보드 버퍼 지우기
			stdin.ReadString('\n')
			continue
		}
		fmt.Printf("입력하신 숫자는 %d입니다\n", number)
		if number%2 == 0 {
			break // 짝수인지 확인
		}
	}
	fmt.Println("for문이 종료되었습니다")
}
