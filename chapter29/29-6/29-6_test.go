package main

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestJsonHandler(t *testing.T) {
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/student", nil)

	mux := MakeWebHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
	student := new(Student)
	// JSON 을 디코딩하여 Student 객체로 변환
	err := json.NewDecoder(res.Body).Decode(student)
	// 에러와 데이터 확인
	assert.Nil(t, err)
	assert.Equal(t, "kim", student.Name)
	assert.Equal(t, 26, student.Age)
	assert.Equal(t, 77, student.Score)
}
