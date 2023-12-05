package main

import "fmt"

func main() {
	var age = 22

	// 논리 연산자 (&&, ||)
	if age >= 10 && age <= 15 {
		fmt.Println("You are young")
	} else if age > 30 || age < 20 {
		fmt.Println("You are not 20s")
	} else {
		fmt.Println("Best age")
	}
}
