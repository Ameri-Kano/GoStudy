package main

import (
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// httptest 패키지를 이용해 http 요청 테스트
func TestIndexHandler(t *testing.T) {
	// 요청 초기화 및 테스트하고 싶은 요청
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)

	mux := MakeWebHandler()
	mux.ServeHTTP(res, req)
	// 요청이 200 OK 로 응답하며 응답이 올바른지 확인
	assert.Equal(t, http.StatusOK, res.Code)
	data, _ := io.ReadAll(res.Body)
	assert.Equal(t, "Hello World!", string(data))
}

func TestBarHandler(t *testing.T) {
	// 요청 초기화 및 테스트하고 싶은 요청
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/bar", nil)

	mux := MakeWebHandler()
	mux.ServeHTTP(res, req)
	// 요청이 200 OK 로 응답하며 응답이 올바른지 확인
	assert.Equal(t, http.StatusOK, res.Code)
	data, _ := io.ReadAll(res.Body)
	assert.Equal(t, "Hello Bar", string(data))
}
