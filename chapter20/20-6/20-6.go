package main

import "fmt"

type Attacker interface {
	Attack()
}

type Monster struct {
	Lv int
}

func (m *Monster) Attack() {
	fmt.Println("Monster Attack")
}

func DoAttack(att Attacker) {
	if att != nil {
		att.Attack()

		var monster *Monster
		// .() 을 통해 인터페이스를 다른 타입이나 다른 인터페이스로 변환
		// 확실하게 알고 있을때 변환해야 한다. 아닌 경우 에러 발생
		monster = att.(*Monster)
		fmt.Println(monster.Lv)
	}
}

func main() {
	// 인터페이스에 아무것도 설정해주지 않으면 기본값은 nil이 된다
	// 문법적으로 에러는 없으나 실행 중 에러(런타임 에러)가 발생
	// nil이므로 비정상적인 메모리에 접근 불가능한 에러 발생
	// var att Attacker
	// att.Attack()

	DoAttack(&Monster{20})
}
