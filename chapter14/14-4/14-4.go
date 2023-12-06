package main

import "fmt"

type Data struct {
	value int
	data  [200]int
}

// ChangeData 함수에서 포인터를 사용하도록 바꿈
// arg는 포인터 변수이므로 (*arg).value 라고 써야 맞겠지만
// Go 에선 아래처럼 작성해도 문제 없이 동작함
func ChangeData(arg *Data) {
	arg.value = 999
	arg.data[100] = 999
}

func main() {
	var data Data

	// 함수에 Data의 주소를 복사하면 값을 변경할 수 있음!
	// 주소 값만 복사하므로 사용하는 메모리 크기를 줄일 수 있음
	ChangeData(&data)
	fmt.Printf("value = %d\n", data.value)
	fmt.Printf("data[100] = %d\n", data.data[100])
}
