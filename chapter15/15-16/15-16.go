package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var str string = "Hello World"
	// 문자열은 불변이므로 슬라이스도 새로운 공간을 만들어서 생성됨
	var slice []byte = []byte(str)

	// 문자열, 슬라이스의 주소 확인을 위해 강제 변환
	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&str))
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&slice))

	// 다른 주소를 가지게 됨을 확인 가능
	fmt.Printf("str:\t%x\n", stringHeader.Data)
	fmt.Printf("slice:\t%x\n", sliceHeader.Data)
}
