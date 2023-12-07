package main

import (
	"fmt"
	"sync"
	"time"
)

type job interface {
	Do()
}

type SquareJob struct {
	index int
}

func (j *SquareJob) Do() {
	fmt.Printf("%d 작업 시작\n", j.index)
	time.Sleep(1 * time.Second)
	fmt.Printf("%d 작업 완료 - 결과 : %d\n", j.index, j.index*j.index)
}

func main() {
	// 각 고루틴이 서로 다른 자원에 접근하게 하는 방법 1
	// job 을 배열에 할당하고 각각 작업을 고루틴으로 실행
	var jobList [10]job

	for i := 0; i < 10; i++ {
		jobList[i] = &SquareJob{i}
	}

	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		job := jobList[i]
		go func() {
			job.Do()
			wg.Done()
		}()
	}
	wg.Wait()
}
