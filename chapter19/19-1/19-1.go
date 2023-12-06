package main

import "fmt"

type account struct {
	balance int
}

// 일반적인 함수 선언
func withdrawFunc(a *account, amount int) {
	a.balance -= amount
}

// 리시버가 지정되어 있는 메서드
func (a *account) withdrawMethod(amount int) {
	a.balance -= amount
}

func main() {
	a := &account{100} // *account

	withdrawFunc(a, 30)

	// a에 속한 메서드
	a.withdrawMethod(30)

	// 메서드와 함수는 기능은 같지만 어떤 타입에 속해있냐 아니냐의 차이
	fmt.Printf("%d \n", a.balance)
}
