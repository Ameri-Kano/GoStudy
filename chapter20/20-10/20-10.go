package main

import "fmt"

type Reader interface {
	Read()
}

type Closer interface {
	Close()
}

type File struct {
}

func (f *File) Close() {
	fmt.Println("Close()")
}

func (f *File) Read() {
	fmt.Println("Read()")
}

func ReadFile(reader Reader) {
	// 타입 변환 시 변수 두개 설정 가능(ok는 성공 여부 bool)
	// 실패하면 기본값, false 반환 Close() 가 없는 경우
	// c, ok := reader.(Closer)
	// fmt.Println(c, ok)
	// 실패하면 실행되지 않을 것이므로 에러를 방지할 수 있음
	// if ok {
	// 	 c.Close()
	// }

	// 위 코드를 더 간단하게 줄일 수 있음
	if c, ok := reader.(Closer); ok {
		c.Close()
	}
}

func main() {
	file := &File{}
	// 만약 Close() 가 없다면 Closer 가 될 수 없으므로 에러 발생
	ReadFile(file)
}
