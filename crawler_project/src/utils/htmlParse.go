package utils

import (
	"os"
	"io/ioutil"
	"fmt"
)

func readHtmlFile(path string)(b []byte)  {
	file, e := os.Open(path)
	defer file.Close()
	if e !=nil {
		fmt.Println("html文件打开失败",e)
	}
	bytes, e2 := ioutil.ReadAll(file)
	if e2!=nil {
		fmt.Printf("文件读取错误",e2)
	}
	return bytes
}

func
