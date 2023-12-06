package main

import (
	"fmt"
	"strings"
)

// ToUpper1 : 문자 철자를 전부 대문자로 변환해줌
// 더할때마다 문자열을 계속 새로운 메모리 공간에 생성하여 메모리 낭비
func ToUpper1(str string) string {
	var rst string
	for _, c := range str {
		if c >= 'a' && c <= 'z' {
			rst += string('A' + (c - 'a'))
		} else {
			rst += string(c)
		}
	}
	return rst
}

// ToUpper2 : ToUpper1과 기능은 같은 기능
// strings.Builder 객체를 이용하면 내부 슬라이스에 저장해 공간 낭비를 없앨 수 있음
func ToUpper2(str string) string {
	var builder strings.Builder
	for _, c := range str {
		if c >= 'a' && c <= 'z' {
			builder.WriteRune('A' + (c - 'a'))
		} else {
			builder.WriteRune(c)
		}
	}
	return builder.String()
}

func main() {
	var str string = "Hello World"

	fmt.Println(ToUpper1(str))
	fmt.Println(ToUpper2(str))
	// strings 패키지에도 구현되어 있음
	fmt.Println(strings.ToUpper(str))
}
