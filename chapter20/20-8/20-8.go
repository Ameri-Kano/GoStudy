package main

import "fmt"

type Stringer interface {
	String() string
}

type Student struct {
}

func (s *Student) String() string {
	return "Student"
}

func main() {
	// Student 가 String() 을 구현하지 않았음
	// 컴파일 에러 발생
	// var stringer Stringer
	// stringer.(*Student)

	// 이것을 빌드할 때에는 확인할 수 없음
	// 그러나 실행하면 nil 이 되므로 에러 발생
	var stringer Stringer
	s := stringer.(*Student)
	fmt.Println(s)
}
