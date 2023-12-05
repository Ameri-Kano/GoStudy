package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	Age   int32   // 4byte
	Score float64 // 8byte
}

func main() {
	user := User{23, 77.2}
	// unsafe.Sizeof() : 변수가 차지하는 메모리 크기 반환(byte 단위)
	// 두 변수를 합치면 12일텐데 왜 16이 나오는가?
	// 에 대한 설명은 메모리 정렬.md 파일에 있음
	fmt.Println(unsafe.Sizeof(user))
}
