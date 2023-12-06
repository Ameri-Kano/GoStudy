package main

import (
	"fmt"
	"time"
)

func main() {
	// 시간 표현 양식에 맞춰서 Time 구조체에 저장 가능
	location, _ := time.LoadLocation("Asia/Seoul")
	const LongForm = "Jan 2, 2006 at 3:04pm"
	t1, _ := time.ParseInLocation(LongForm, "Dec 6, 2023 at 3:00pm", location)
	fmt.Println(t1, t1.Location(), t1.UTC())

	const shortForm = "2006-Jan-02"
	t2, _ := time.Parse(shortForm, "2023-Dec-06")
	fmt.Println(t2, t2.Location())

	// 형식이 바르지 않으면 에러를 반환하며 시간이 올바르게 저장되지 않음
	t3, err := time.Parse("2021-06-01 15:20:21", "2023-12-06 15:21:00")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t3, t3.Location())

	// t2 - t1 의 결과(시간)
	d := t2.Sub(t1)
	fmt.Println(d)
}
