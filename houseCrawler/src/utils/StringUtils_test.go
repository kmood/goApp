package utils

import (
	"strings"
	"testing"
)

func Test_DelInvisibleChar(t *testing.T) {
	char := DelInvisibleChar("tests/  test")
	if strings.EqualFold(char, "tests/test") {
		t.Log("测试通过")
	}
}
