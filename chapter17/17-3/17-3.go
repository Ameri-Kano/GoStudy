package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

var stdin = bufio.NewReader(os.Stdin)

func InputIntValue() (int, error) {
	var n int
	_, err := fmt.Scanln(&n)
	if err != nil {
		// 잘못된 입력 발생시 버퍼 비워주기
		stdin.ReadString('\n')
	}
	return n, err
}

func main() {
	// 시드 설정해줄 필요 없음
	// rand.Seed(time.Now().UnixNano())

	r := rand.Intn(100)
	cnt := 1
	for {
		fmt.Println("숫자를 입력하세요")
		n, err := InputIntValue()
		if err != nil {
			fmt.Println("숫자만 입력하세요")
		} else {
			if n > r {
				fmt.Println("입력한 숫자가 더 큽니다")
			} else if n < r {
				fmt.Println("입력한 숫자가 더 작습니다")
			} else {
				fmt.Println("정답입니다, 축하합니다! 시도한 횟수:", cnt)
				break
			}
			cnt++
		}
	}
}
