package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	A int8 // 1
	B int  // 8
	C int8 // 1
	D int  // 8
	E int8 // 1
}

func main() {
	user := User{1, 2, 3, 4, 5}
	// 결과는 40이 나옴 (필드의 총합은 19)
	fmt.Println(unsafe.Sizeof(user))
}
