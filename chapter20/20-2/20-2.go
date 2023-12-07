package main

import (
	"GoStudyProject/chapter20/koreaPost"
)

type Sender interface {
	Send(parcel string)
}

func SendBook(name string, sender Sender) {
	sender.Send(name)
}

func main() {
	// 사용하는 환경에 따라 인터페이스가 받을 타입만 정해주면
	// 코드를 복잡하게 수정하지 않아도 됨
	// 추상화를 통해 복잡하게 구현된 각각의 코드를 알지 않아도 됨
	// var sender Sender = &fedex.FedexSender{}
	var sender Sender = &koreaPost.PostSender{}
	SendBook("어린 왕자", sender)
	SendBook("그리스인 조르바", sender)
}
