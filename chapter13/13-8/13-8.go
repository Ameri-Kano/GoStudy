package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	A int8 // 1
	C int8 // 1
	E int8 // 1
	D int  // 8
	B int  // 8
}

func main() {
	user := User{1, 2, 3, 4, 5}
	// 필드의 순서를 조정하여 결과가 24가 됨
	fmt.Println(unsafe.Sizeof(user))
}
