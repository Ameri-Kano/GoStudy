package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func barHandler(w http.ResponseWriter, r *http.Request) {
	// 쿼리 인수를 가져와서
	values := r.URL.Query()
	// name 이라는 키에 값이 있는지 확인
	name := values.Get("name")
	if name == "" {
		name = "World"
	}
	// id 값을 가져와서 정수로 변환
	id, _ := strconv.Atoi(values.Get("id"))
	fmt.Fprintf(w, "Hello %s! id:%d", name, id)
}

func main() {
	// localhost:3000/bar?name=kim&id=13 요청시
	// Hello kim! id:13 응답
	http.HandleFunc("/bar", barHandler)
	http.ListenAndServe(":3000", nil)
}
