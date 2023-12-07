package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	rand.Seed(time.Now().UnixNano())

	wg.Add(2)
	fork := &sync.Mutex{}
	spoon := &sync.Mutex{}

	// mutex 를 잘못 사용하면 데드락이 발생할 수 있음!
	// 서로 똑같이 수저, 포크 순서대로 집게 하면 이 코드에서는 데드락을 막을 수 있음
	// 그러나 모든 경우를 해결할 수는 없음
	go diningProblem("A", fork, spoon, "포크", "수저")
	go diningProblem("B", spoon, fork, "수저", "포크")
	wg.Wait()
}

func diningProblem(name string, first, second *sync.Mutex, firstname, secondName string) {
	for i := 0; i < 100; i++ {
		fmt.Printf("%s 밥을 먹으려 합니다\n", name)
		first.Lock()
		fmt.Printf("%s %s 획득\n", name, firstname)
		second.Lock()
		fmt.Printf("%s %s 획득\n", name, secondName)

		fmt.Printf("%s 밥을 먹습니다\n", name)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

		second.Unlock()
		first.Unlock()
	}
	wg.Done()
}
