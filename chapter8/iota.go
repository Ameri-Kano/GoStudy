package main

import "fmt"

// iota 를 이용하여 간단하게 열거형 상수 정의 가능
// iota = 0, 1, 2, 3...
const (
	MasterRoom uint8 = 1 << iota
	LivingRoom
	BathRoom
	SmallRoom
)

func SetLight(rooms, room uint8) uint8 {
	return rooms | room
}

// ResetLight 는 방에 해당하는 비트에 비트클리어 연산 수행
func ResetLight(rooms, room uint8) uint8 {
	return rooms &^ room
}

func IsLightOn(rooms, room uint8) bool {
	return rooms&room == room
}

func TurnLights(rooms uint8) {
	if IsLightOn(rooms, MasterRoom) {
		fmt.Println("안방에 불을 켠다")
	}
	if IsLightOn(rooms, LivingRoom) {
		fmt.Println("거실에 불을 켠다")
	}
	if IsLightOn(rooms, BathRoom) {
		fmt.Println("화장실에 불을 켠다")
	}
	if IsLightOn(rooms, SmallRoom) {
		fmt.Println("작은방에 불을 켠다")
	}
}

func main() {
	var rooms uint8 = 0
	rooms = SetLight(rooms, MasterRoom)
	rooms = SetLight(rooms, BathRoom)
	rooms = SetLight(rooms, SmallRoom)
	rooms = ResetLight(rooms, SmallRoom)

	TurnLights(rooms)
}
