package main

import (
	"GoStudyProject/usepkg/custompkg"
	"GoStudyProject/usepkg/exinit"
	"fmt"
	"github.com/guptarohit/asciigraph"
	"github.com/tuckersGo/musthaveGo/ch16/expkg"
)

func main() {
	custompkg.PrintCustom()

	// 대문자로 시작하는 변수, 상수는 외부로 공개되어 사용 가능(custompkg 에서)
	s := custompkg.Student{}
	s.Name = "화랑"
	s.Age = 32
	// s.score = 100	// 공개되어있지 않으므로 외부에서 사용할 수 없음
	fmt.Println(s.Name, s.Age)
	fmt.Println(custompkg.PI)
	custompkg.Ratio = 10
	fmt.Println(custompkg.Ratio)

	expkg.PrintSample()

	// 외부 패키지를 가져와서 사용
	// go mod init GoStudyProject/usepkg 으로 새 모듈 생성
	// go mod tidy 로 필요한 외부 모듈 다운로드
	// 다운받은 외부 모듈은 GOPATH 에 저장되어 이후에 다른 곳에서도 사용 가능
	data := []float64{3, 4, 5, 6, 9, 7, 5, 8, 5, 10, 2, 7, 2, 5, 6}
	graph := asciigraph.Plot(data)
	fmt.Println(graph)

	// 함수를 실행하기 전 패키지의 init() 이 제일 먼저 실행된다 (패키지 초기화)
	// init() 은 오직 한번만 실행된다
	exinit.PrintD()
}
