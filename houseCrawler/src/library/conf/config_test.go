package conf

import (
	"fmt"
	"testing"
)
//单元测试关键点 文件名为xxx_test.go  函数以Test为前缀
func TestGetConfig(t *testing.T) {
	conf := GetConfig()
	t.Log(fmt.Sprintf("%v-%v",conf.Log,conf.Mysql))
}
