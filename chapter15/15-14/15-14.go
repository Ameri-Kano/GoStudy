package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {

	str1 := "Hello 월드"
	str2 := str1

	// string 을 복사할 때 문자열 값 전체를 복사하지 않는다
	// StringHeader 구조체는 문자열의 주소와 길이(전체 용량)을 가진다
	stringHeader1 := (*reflect.StringHeader)(unsafe.Pointer(&str1))
	stringHeader2 := (*reflect.StringHeader)(unsafe.Pointer(&str2))

	// 두 string 의 data 가 같음을 확인할 수 있다
	fmt.Println(stringHeader1)
	fmt.Println(stringHeader2)
}
