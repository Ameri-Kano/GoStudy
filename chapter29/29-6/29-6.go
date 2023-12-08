package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Student struct {
	Name  string
	Age   int
	Score int
}

func MakeWebHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/student", StudentHandler)
	return mux
}

func StudentHandler(w http.ResponseWriter, r *http.Request) {
	student := Student{"kim", 26, 77}
	// JSON 형태로 변환
	data, _ := json.Marshal(student)
	// 헤더에 JSON 형태임을 명시
	w.Header().Add("content-type", "application/json")
	// 응답 메시지가 OK임을 표시
	w.WriteHeader(http.StatusOK)
	// 결과 전송
	fmt.Fprint(w, string(data))
}

func main() {
	http.ListenAndServe(":3000", MakeWebHandler())
}
