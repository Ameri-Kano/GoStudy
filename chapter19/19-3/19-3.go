package main

import "fmt"

type account struct {
	balance   int
	firstname string
	lastname  string
}

func (a1 *account) withdrawPointer(amount int) {
	a1.balance -= amount
}

func (a2 account) withdrawValue(amount int) {
	a2.balance -= amount
}

func (a3 account) withdrawValue2(amount int) account {
	a3.balance -= amount
	return a3
}

func main() {
	var mainA *account = &account{100, "joe", "park"}

	mainA.withdrawPointer(30)
	fmt.Println(mainA.balance) // 70

	// 단순히 값이 복사되므로 영향을 주지 않음
	mainA.withdrawValue(20)
	fmt.Println(mainA.balance) // 70

	// 값을 변경하여 반환해주면 정상 작동
	*mainA = mainA.withdrawValue2(20)
	fmt.Println(mainA.balance) // 50
}
