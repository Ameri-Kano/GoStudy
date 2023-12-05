package main

import "fmt"

// ColorType 은 int 와 같지만 이름을 새롭게 지어줌(별칭)
type ColorType int

const (
	Red ColorType = iota
	Blue
	Green
	Yellow
)

// colorToString 은 색깔 열거값 상수에 따른 문자열 반환
func colorToString(color ColorType) string {
	switch color {
	case Red:
		return "Red"
	case Blue:
		return "Blue"
	case Green:
		return "Green"
	case Yellow:
		return "Yellow"
	default:
		return "Undefined"
	}
}

func getMyFavoriteColor() ColorType {
	return Red
}

func main() {
	// const 열거값과 switch 문의 활용
	fmt.Println("My favorite color is", colorToString(getMyFavoriteColor()))
}
