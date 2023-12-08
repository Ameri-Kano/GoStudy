package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Test 코드 작성 방법
// 테스트 코드 파일명은 _test.go 로 끝나야 한다
// import "testing" 을 포함해야 한다
// func TestXxxx(t *testing.T) 형태이여야 한다
func TestSquare1(t *testing.T) {
	// 외부 라이브러리를 이용해 테스트 코드를 더욱 간단하게 작성할 수 있음
	assert.Equal(t, 81, square(9), "square(9) should be 81")
	// rst := square(9)
	// if rst != 81 {
	//	 t.Errorf("square(9) should be 81 but returns %d", rst)
	// }
}

func TestSquare2(t *testing.T) {
	assert.Equal(t, 9, square(3), "square(3) should be 9")
	assert.Equal(t, 0, square(0), "square(9) should be 0")
}
