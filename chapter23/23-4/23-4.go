package main

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func MultipleFromString(str string) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(str))
	scanner.Split(bufio.ScanWords)

	pos := 0
	a, n, err := readNextInt(scanner)
	if err != nil {
		// %w 를 이용하여 에러 래핑
		return 0, fmt.Errorf("Failed to readNextInt(), pos:%d err:%w", pos, err)
	}

	pos += n + 1
	b, n, err := readNextInt(scanner)
	if err != nil {
		return 0, fmt.Errorf("Failed to readNextInt(), pos:%d err:%w", pos, err)
	}
	return a * b, nil
}

// 단어를 읽어서 숫자로 변환하여 반환
// 변환된 숫자, 읽은 글자 수, 에러를 반환함
func readNextInt(scanner *bufio.Scanner) (int, int, error) {
	if !scanner.Scan() {
		return 0, 0, fmt.Errorf("Failed to scan")
	}
	word := scanner.Text()
	number, err := strconv.Atoi(word)
	if err != nil {
		return 0, 0, fmt.Errorf("Failed to convert word to int, word:%s err:%w",
			word, err)
	}
	return number, len(word), nil
}

func readEq(eq string) {
	rst, err := MultipleFromString(eq)
	if err == nil {
		fmt.Println(rst)
	} else {
		fmt.Println(err)
		var numError *strconv.NumError
		// 감싸진 에러들 중 numError 가 포함되어 있다면 그 에러를 변환
		if errors.As(err, &numError) {
			fmt.Println("NumberError", numError)
		}
	}
}

func main() {
	// 에러 래핑 예제
	// 여러 가지의 에러를 감싸서 각 에러 타입별로 처리 가능
	readEq("123 3")
	readEq("123 abc")
}
