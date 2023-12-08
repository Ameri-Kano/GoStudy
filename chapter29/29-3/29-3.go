package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 새로운 Servemux 에 핸들러를 등록함
	// 기본 핸들러를 사용하지 않으므로 핸들러에 다양한 기능을 추가할 수 있음
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})
	mux.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello Bar")
	})

	http.ListenAndServe(":3000", mux)
}
