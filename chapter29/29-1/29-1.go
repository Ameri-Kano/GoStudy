package main

import (
	"fmt"
	"net/http"
)

// 웹 서버 만들기 기본 코드
func main() {
	// "/" 경로에 해당하는 요청 수신 핸들러 함수
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World!")
	})

	// 서버 포트를 3000번으로 설정, 핸들러 nil 설정시 기본값 (DefaultServeMux) 사용
	// localhost:3000 접속시 Hello World!가 화면에 출력됨
	http.ListenAndServe(":3000", nil)
}
