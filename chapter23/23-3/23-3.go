package main

import "fmt"

type PasswordError struct {
	Len        int
	RequireLen int
}

func (err PasswordError) Error() string {
	return "암호 길이가 짧습니다."
}

func RegisterAccount(name, password string) error {
	if len(password) < 8 {
		// Error() 메소드가 구현되어 있으므로 Error 객체가 될 수 있음
		return PasswordError{len(password), 8}
	}
	// ...
	return nil
}

func main() {
	err := RegisterAccount("myId", "myPw")

	if err != nil {
		if errInfo, ok := err.(PasswordError); ok {
			fmt.Printf("%v Len:%d RequireLen:%d\n",
				errInfo, errInfo.Len, errInfo.RequireLen)
		}
	} else {
		fmt.Println("회원 가입 완료")
	}
}
