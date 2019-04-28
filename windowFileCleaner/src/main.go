package main

import (
	"fmt"
	"github.com/lxn/walk"
	"github.com/lxn/walk/declarative"
	"log"
	"path/filepath"
)

func main() {

	fileInfoList := NewFileInfoList()
	var mainWindow *walk.MainWindow
	var lineText *walk.LineEdit
	var tableView *walk.TableView
	var vsSplitter *walk.Splitter
	var hsSplitter *walk.Splitter
	var webView *walk.WebView
	if _, err := (declarative.MainWindow{
		MinSize:  declarative.Size{600, 400},
		Size:     declarative.Size{1024, 640},
		Title:    "目录",
		AssignTo: &mainWindow,
		Layout:   declarative.HBox{MarginsZero: true},
		Children: []declarative.Widget{
			declarative.VSplitter{

				AssignTo: &vsSplitter,
				Children: []declarative.Widget{
					declarative.Label{
						Text: "输入路径:",
					},
					declarative.LineEdit{
						AssignTo: &lineText,
						OnKeyUp: func(key walk.Key) {
							if key == walk.KeyReturn {
								fileInfoList.SetFileInfoList(lineText.Text())
								fmt.Println(fileInfoList.TotalSize / 1024 / 1024)
							}
						},
					},
					declarative.HSplitter{
						AssignTo: &hsSplitter,

						Children: []declarative.Widget{
							declarative.TableView{
								AssignTo:      &tableView,
								StretchFactor: 2,
								Columns: []declarative.TableViewColumn{
									declarative.TableViewColumn{
										DataMember: "Name",
										Width:      192,
									},
									declarative.TableViewColumn{
										DataMember: "Size",
										Format:     "%d",
										Alignment:  declarative.AlignFar,
										Width:      64,
									},
									declarative.TableViewColumn{
										DataMember: "IsDir",
										Width:      120,
									},
									declarative.TableViewColumn{
										DataMember: "Zt",
										Width:      120,
									},
								},
								Model: fileInfoList,

								OnCurrentIndexChanged: func() {
									//鼠标点击选中时触发
									var url string
									if index := tableView.CurrentIndex(); index > -1 {
										name := fileInfoList.items[index].Name
										dir := lineText.Text()
										url = filepath.Join(dir, "/", name)
									}

									webView.SetURL(url)

								},
								OnItemActivated: func() {
									var url string
									//双击时的事件
									if index := tableView.CurrentIndex(); index > -1 {
										name := fileInfoList.items[index].Name
										dir := lineText.Text()
										url = filepath.Join(dir, "/", name)
									}
									lineText.SetText(url)
									fileInfoList.SetFileInfoList(url)
								},
							},
							declarative.WebView{
								AssignTo:      &webView,
								StretchFactor: 2,
							},
						},
					},
				},
			},
		},
	}.Run()); err != nil {
		log.Fatal(err)
	}
}
