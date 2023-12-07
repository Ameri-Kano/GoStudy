package main

import "fmt"

// Stringer 인터페이스
// type 이름 interface
type Stringer interface {
	String() string
}

type Student struct {
	Name string
	Age  int
}

func (s Student) String() string {
	return fmt.Sprintf("안녕! 나는 %d살 %s라고 해", s.Age, s.Name)
}

func main() {
	student := Student{"철수", 12}
	var stringer Stringer

	// 인터페이스가 가지고 있는 메서드만 쓸 수 있다
	// student.String() 이라 해도 같을텐데
	// 인터페이스는 왜 사용할까?
	stringer = student
	fmt.Printf("%s\n", stringer.String())
}
