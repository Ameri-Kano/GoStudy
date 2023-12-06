package custompkg

import "fmt"

type Student struct {
	Name  string
	Age   int
	score int
}

var Ratio int
var ttt int // 소문자로 시작하는 상수, 변수, 함수는 외부로 공개되지 않음

const PI = 3.14
const pi2 = 3.1415

func PrintCustom() {
	fmt.Println("This is custom package!")
}

func printcustom2() {
	fmt.Println("This is custom package 2!")
}
