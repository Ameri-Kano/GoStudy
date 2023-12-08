package main

import (
	"fmt"
	"sync"
	"time"
)

type Car struct {
	Body  string
	Tire  string
	Color string
}

var wg sync.WaitGroup
var startTime = time.Now()

func main() {
	// 채널을 이용하여 생산자/소비자 패턴 구현
	tireCh := make(chan *Car)
	paintCh := make(chan *Car)

	fmt.Printf("Start Factory\n")

	// 각각의 역할을 하는 3개의 함수(바디, 타이어, 페인팅)
	wg.Add(3)
	go MakeBody(tireCh)
	go InstallTire(tireCh, paintCh)
	go PaintCar(paintCh)

	wg.Wait()
	fmt.Println("Close the factory")
}

// MakeBody 는 10초 동안 1초에 한 번씩 바디를 만들어 타이어 채널로 전달
func MakeBody(tireCh chan *Car) {
	tick := time.Tick(time.Second)
	after := time.After(10 * time.Second)
	for {
		select {
		case <-tick:
			car := &Car{}
			car.Body = "Sports car"
			tireCh <- car
		case <-after:
			close(tireCh)
			wg.Done()
			return
		}
	}
}

// InstallTire 는 전달받은 바디가 만들어진 차에 타이어 설치하여 페인트 채널로 전달
func InstallTire(tireCh, paintCh chan *Car) {
	for car := range tireCh {
		time.Sleep(time.Second)
		car.Tire = "Winter tire"
		paintCh <- car
	}
	wg.Done()
	close(paintCh)
}

// PaintCar 는 페인트칠 이후 완성된 차의 정보 출력
func PaintCar(paintCh chan *Car) {
	// 데이터가 전부 처리되면 채널이 닫히므로 프로그램 정상 종료
	for car := range paintCh {
		time.Sleep(time.Second)
		car.Color = "Red"
		// 지금 시간 - 시작 시간
		duration := time.Now().Sub(startTime)
		fmt.Printf("%.2f Complete Car: %s %s %s\n", duration.Seconds(),
			car.Body, car.Tire, car.Color)
	}
	wg.Done()
}
