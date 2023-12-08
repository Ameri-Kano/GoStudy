package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})

	// HTTPS 서버 시작
	err := http.ListenAndServeTLS(":3000", "localhost.crt", "localhost.key", nil)
	// 오류 발생하면 로그 출력
	if err != nil {
		log.Fatal(err)
	}
}
