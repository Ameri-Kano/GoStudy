package main

import "fmt"

// if 초기문; 조건문 {
//	문장
// }

// 아래 코드는 UploadFile()을 먼저 실행하고 success의 값을 확인함
// if filename, success := UploadFile(); success {
//	 fmt.Println("Upload success", filename)
// } else {
//	 fmt.Println("Failed to upload")
// }

func getMyAge() (int, bool) {
	return 25, true
}

func main() {

	if age, ok := getMyAge(); ok && age < 20 {
		fmt.Println("You are young", age)
	} else if ok && age < 30 {
		fmt.Println("Nice age", age)
	} else if ok {
		fmt.Println("You are beautiful", age)
	} else {
		fmt.Println("Error")
	}

	// if문을 벗어났으므로 age는 사라져 사용할 수 없음
	// 아래 코드는 에러
	// fmt.Println("Your age is", age)
}
