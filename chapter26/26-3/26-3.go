package main

import (
	"bufio"
	"fmt"
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

func GetFileList(path string) ([]string, error) {
	return filepath.Glob(path)
}

func FindWordInAllFiles(word, path string) []FindInfo {
	findInfos := []FindInfo{}

	fileList, err := GetFileList(path)
	if err != nil {
		fmt.Println("파일 경로가 잘못되었습니다. err:", err, "path:", path)
		return findInfos
	}
	for _, filename := range fileList {
		// 찾은 FindInfo 에 순서대로 추가
		findInfos = append(findInfos, FindWordInFile(word, filename))
	}
	return findInfos
}

func FindWordInFile(word, filename string) FindInfo {
	findInfo := FindInfo{filename, []LineInfo{}}
	file, err := os.Open(filename)

	if err != nil {
		fmt.Println("파일을 찾을 수 없습니다.", filename)
		return findInfo
	}
	defer file.Close()

	lineNo := 1
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// 한 줄씩 읽어서
		line := scanner.Text()
		// 해당 단어를 포함하면
		if strings.Contains(line, word) {
			// LineInfo에 추가
			findInfo.lines = append(findInfo.lines, LineInfo{lineNo, line})
		}
		lineNo++
	}
	return findInfo
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("2개 이상의 실행 인수가 필요합니다. ex) 26-1 word filepath")
		return
	}

	word := os.Args[1]
	files := os.Args[2:]
	findInfos := []FindInfo{}
	for _, path := range files {
		findInfos = append(findInfos, FindWordInAllFiles(word, path)...)
	}
	for _, findInfo := range findInfos {
		fmt.Println(findInfo.filename)
		fmt.Println("------------------------------------------------")
		for _, lineInfo := range findInfo.lines {
			fmt.Println("\t", lineInfo.lineNo, "\t", lineInfo.line)
		}
		fmt.Println("------------------------------------------------")
	}
}
