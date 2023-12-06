package main

import "fmt"

type Data struct {
	value int
	data  [200]int
}

func ChangeData(arg Data) {
	arg.value = 999
	arg.data[100] = 999
}

func main() {
	var data Data

	// 함수에는 단순히 구조체의 값이 복사된다
	// 그러므로 ChangeData 함수의 arg와 main 함수의 data의 공간이 다름
	// 이를 해결하기 위하여 포인터를 사용할 수 있다(14-4)
	ChangeData(data)
	// 값을 바꾸려 했는데 결과적으로 둘 다 0이 나옴
	fmt.Printf("value = %d\n", data.value)
	fmt.Printf("data[100] = %d\n", data.data[100])
}
