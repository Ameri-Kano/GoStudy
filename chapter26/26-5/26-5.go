package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type LineInfo struct {
	lineNo int
	line   string
}

type FindInfo struct {
	filename string
	lines    []LineInfo
}

// GetFileList 를 모든 하위 폴더까지 모두 포함해서 검색할 수 있도록 변경
func GetFileList(pattern string) ([]string, error) {
	fileList := []string{}

	// filepath.Walk() : 기준 폴더에서부터 아래의 모든 파일을 순회
	// . 이 현재 디렉토리이므로 현재 디렉토리 아래의 모든 디렉토리를 포함하게 됨
	err := filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
		// info 를 확인해서 디렉토리가 아니면(파일이면)
		if !info.IsDir() {
			// 패턴과 일치하는지 확인
			matched, _ := filepath.Match(pattern, info.Name())
			// 패턴과 일치하면 파일 리스트에 추가
			if matched {
				fileList = append(fileList, path)
			}
		}
		return nil
	})
	if err != nil {
		return []string{}, err
	}
	return fileList, nil
}

func FindWordInAllFiles(word, path string) (int, chan FindInfo) {
	filelist, err := GetFileList(path) // 실행인자 가져오기
	if err != nil {
		fmt.Println("파일을 찾을 수 없습니다. err:", err)
		return 0, nil
	}

	ch := make(chan FindInfo)
	cnt := len(filelist)
	for _, filename := range filelist {
		go FindWordInFile(word, filename, ch)
	}
	return cnt, ch
}

func FindWordInFile(word, filename string, ch chan FindInfo) {
	findInfo := FindInfo{filename, []LineInfo{}}
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("파일을 찾을 수 없습니다. ", filename)
		ch <- findInfo
		return
	}
	defer file.Close()

	lineNo := 1
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, word) {
			findInfo.lines = append(findInfo.lines, LineInfo{lineNo, line})
		}
		lineNo++
	}
	ch <- findInfo
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("2개 이상의 실행 인수가 필요합니다. ex) 26-5 word filepath")
		return
	}

	word := os.Args[1]
	path := os.Args[2]
	cnt, ch := FindWordInAllFiles(word, path)
	recvCnt := 0
	for findInfo := range ch {
		fmt.Println(findInfo.filename)
		fmt.Println("--------------------------------")
		for _, lineInfo := range findInfo.lines {
			fmt.Println("\t", lineInfo.lineNo, "\t", lineInfo.line)
		}
		fmt.Println("--------------------------------")
		fmt.Println()
		recvCnt++
		if recvCnt == cnt {
			// all received
			break
		}
	}
}
