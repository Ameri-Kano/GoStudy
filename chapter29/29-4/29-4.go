package main

import "net/http"

func main() {
	// "/static" 으로 요청이 오지만 제거해주면 특정 경로의 파일을 올바르게 찾아줄 수 있음
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":3000", nil)
}
