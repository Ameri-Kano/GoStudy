package main

import (
	"fmt"
	"sync"
	"time"
)

type Account struct {
	Balance int
}

func main() {
	var wg sync.WaitGroup

	account := &Account{0}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			for {
				// 하나의 balance 에 끊임없이 접근하므로 에러가 발생할 수 있음
				// 이러한 문제를 동시성 문제라고 함
				DepositAndWithdraw(account)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func DepositAndWithdraw(account *Account) {
	if account.Balance < 0 {
		panic(fmt.Sprintf("Balance should not be negative value: %d",
			account.Balance))
	}
	account.Balance += 1000
	time.Sleep(time.Millisecond)
	account.Balance -= 1000
}
