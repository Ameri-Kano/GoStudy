package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func (u User) String() string {
	return fmt.Sprintf("%s %d", u.Name, u.Age)
}

type VIPUser struct {
	User     // embedded field
	VIPLevel int
}

func (v VIPUser) vipLevel() int {
	return v.VIPLevel
}

func main() {
	// String() 은 User 에 정의되어 있으며 VIPUser 가 User 를 포함하므로 정상적으로 동작
	// 상속같은 느낌이 들지만 편의를 위해서지 상속이 아니다
	vip := VIPUser{User{"Hwarang", 34}, 5}
	fmt.Println(vip.String())
}
