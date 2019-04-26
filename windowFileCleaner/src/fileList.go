package main

import (
	"fmt"
	"github.com/lxn/walk"
	"io/ioutil"
	"os"
	"path/filepath"
)

type FileInfo struct {
	Name  string
	IsDir bool
	Size  int64
}

func NewFileInfo() *FileInfo {
	return &FileInfo{}
}

//type FileInfoList struct {
//	FileInfos []*FileInfo
//	Path string
//	TotalSize int64
//}
type FileInfoList struct {
	walk.SortedReflectTableModelBase
	Path      string
	TotalSize int64
	items     []*FileInfo
}

func (m *FileInfoList) Items() interface{} {
	return m.items
}
func (m *FileInfoList) Image(row int) interface{} {
	return filepath.Join(m.Path, m.items[row].Name)
}

//NewFileInfoList
func NewFileInfoList() *FileInfoList {
	return &FileInfoList{}
}
func (f *FileInfoList) SetFileInfoList(path string) {
	file, _ := os.Open(path)
	f.items = nil //清空item
	infos, _ := file.Readdir(0)
	var total int64 = 0
	for _, fileinf := range infos {
		info := NewFileInfo()
		var sizeTem int64 = 0
		if fileinf.IsDir() {
			info.IsDir = true
			sizeTem, _ = getDirContainSize(path + filepath.ToSlash("/") + fileinf.Name())
		} else {
			info.IsDir = false
			sizeTem = fileinf.Size()
		}
		info.Size = sizeTem >> 20
		info.Name = fileinf.Name()
		total += sizeTem
		f.items = append(f.items, info)
	}
	f.Path = path
	f.TotalSize = total
	f.PublishRowsReset()
	defer file.Close()

}

//getDirContainSize  获取文件夹大小
func getDirContainSize(dirPath string) (size int64, err error) {

	var fileSize int64 = 0
	infos, e := ioutil.ReadDir(dirPath)
	if e != nil {
		return 0, nil
	}
	for _, fileInfo := range infos {
		if fileInfo.IsDir() {
			dirPath_ := fmt.Sprint(dirPath, filepath.ToSlash("/"), fileInfo.Name())
			i, err := getDirContainSize(dirPath_)
			if err != nil {
				return 0, nil
			}
			fileSize += i
		} else {
			fileSize += fileInfo.Size()
		}
	}
	return fileSize, nil
}
