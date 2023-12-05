package main

import "fmt"

type User struct {
	Name string
	ID   string
	Age  int
}

type VIPUser struct {
	User     // 내장된 필드
	VIPLevel int
	price    int
}

func main() {
	user := User{"송하나", "hana", 23}
	vip := VIPUser{
		User{"화랑", "hwarang", 40},
		3,
		250,
	}

	// 내장된 구조체 필드의 경우 . 하나만으로 접근 가능
	// 그러나 이름이 겹치는 경우 내장된 필드가 아닌 원래 구조체의 필드를 우선적으로 접근
	fmt.Printf("유저 : %s ID : %s 나이 : %d\n", user.Name, user.ID, user.Age)
	fmt.Printf("VIP 유저: %s ID: %s 나이: %d VIP레벨 : %d VIP가격: %d만원\n",
		vip.Name,
		vip.ID,
		vip.Age,
		vip.VIPLevel,
		vip.price,
	)
}
