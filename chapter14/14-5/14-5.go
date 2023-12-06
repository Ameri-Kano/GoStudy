package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func NewUser(name string, age int) *User {
	var u = User{name, age}
	return &u
}

func main() {

	userPointer := NewUser("AAA", 23)

	// 원래대로라면 u는 이미 없어져서 에러가 발생해야 하지만 에러가 발생하지 않음
	// 탈출 분석(Escape Analyzing) 을 통해 메모리를 컴퓨터의 어디에 만들지 결정
	// 일반적으로 함수 지역 변수는 스택 영역, 함수 외부로 공개되는 메모리는 힙 영역
	// userPointer 변수로 탈출하므로 처음부터 변수를 힙 영역에 할당해 사라지지 않게 함
	// 이를 통해 C나 C++에서 발생할 수 있던 댕글링(dangling, 사라진 메모리를 가리키는 오류) 오류 방지
	fmt.Println(userPointer)
}
