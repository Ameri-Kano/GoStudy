package main

import (
	"encoding/json"
	"net/http"
	"sort"
	"strconv"

	"github.com/gorilla/mux"
)

type Student struct {
	Id    int
	Name  string
	Age   int
	Score int
}

// 학생 목록 저장
var students map[int]Student
var lastId int

func MakeWebHandler() http.Handler {
	mux := mux.NewRouter()
	// /students 로 GET 요청이 들어오면 GetStudentListHandler 함수 실행
	mux.HandleFunc("/students", GetStudentListHandler).Methods("GET")
	// /students/id 로 GET 요청이 들어오면 GetStudentHandler 실행
	mux.HandleFunc("/students/{id:[0-9]+}", GetStudentHandler).Methods("GET")
	// /students 로 POST 요청이 들어오면 PostStudentHandler 실행
	mux.HandleFunc("/students", PostStudentHandler).Methods("POST")
	// /students/id 로 DELETE 요청이 들어오면 DeleteStudentHandler 실행
	mux.HandleFunc("/students/{id:[0-9]+}", DeleteStudentHandler).Methods("DELETE")

	students = make(map[int]Student)
	students[1] = Student{1, "aaa", 16, 87}
	students[2] = Student{2, "bbb", 18, 98}
	lastId = 2

	return mux
}

type Students []Student

func (s Students) Len() int {
	return len(s)
}

func (s Students) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Students) Less(i, j int) bool {
	// Id 를 기준으로 정렬
	return s[i].Id < s[j].Id
}

func GetStudentListHandler(w http.ResponseWriter, r *http.Request) {
	list := make(Students, 0)
	for _, student := range students {
		list = append(list, student)
	}
	sort.Sort(list)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	// JSON 으로 변경하여 결과 쓰기
	json.NewEncoder(w).Encode(list)
}

func GetStudentHandler(w http.ResponseWriter, r *http.Request) {
	// 경로 변수 가져오기
	// 경로를 {id:[0-9]+} 로 설정했으므로 자동으로 id 라고 저장
	// [0-9]+ 는 한 개 이상의 숫자를 의미하는 정규표현식
	vars := mux.Vars(r)
	// string -> int
	id, _ := strconv.Atoi(vars["id"])
	student, ok := students[id]
	// 존재하지 않을 시 에러
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(student)
}

func PostStudentHandler(w http.ResponseWriter, r *http.Request) {
	var student Student
	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	lastId++
	student.Id = lastId
	students[lastId] = student
	w.WriteHeader(http.StatusCreated)
}

func DeleteStudentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	_, ok := students[id]
	if !ok {
		// 없으면 404 에러
		w.WriteHeader(http.StatusNotFound)
		return
	}
	delete(students, id)
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.ListenAndServe(":3000", MakeWebHandler())
}
