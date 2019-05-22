package utils

import "regexp"

//DelInvisibleChar 清除不可见字符
func DelInvisibleChar(str string) string {
	compile := regexp.MustCompile(`[\x00-\x1F | \x7F ]`)
	result := compile.ReplaceAllString(str, "")
	return result
}
