package gui

import (
	"github.com/lxn/walk"
	"github.com/lxn/walk/declarative"
	"log"
	"selenium"
)

type SnapUp struct {
	SnapUpPageUrl, UserName, Password string
}

func NewSnapUp() *SnapUp {
	return &SnapUp{}
}

func SnapUpGui() {
	var mw *walk.MainWindow
	var outTE *walk.TextEdit
	var db *walk.DataBinder

	//var outLog string
	snapUp := NewSnapUp()
	if _, err := (declarative.MainWindow{
		AssignTo: &mw,
		Title:    "HUA WEI Snap up",
		MinSize:  declarative.Size{500, 400},
		Layout:   declarative.VBox{},
		DataBinder: declarative.DataBinder{
			AssignTo:       &db,
			Name:           "snapUp",
			DataSource:     snapUp,
			ErrorPresenter: declarative.ToolTipErrorPresenter{},
		},
		Children: []declarative.Widget{
			declarative.Label{
				Text: "基本信息:",
			},
			declarative.Composite{
				Layout: declarative.Grid{Columns: 2},
				Children: []declarative.Widget{
					declarative.Label{
						Text: "用户名:",
					},
					declarative.LineEdit{
						Text: declarative.Bind("UserName"),
					},

					declarative.Label{
						Text: "密码:",
					},
					declarative.LineEdit{
						Text: declarative.Bind("Password"),
					},

					declarative.Label{
						Text: "抢购页面:",
					},
					declarative.LineEdit{
						Text: declarative.Bind("SnapUpPageUrl"),
					},
				},
			},
			declarative.Label{
				Text: "-------------------------------------",
			},
			declarative.PushButton{
				Text: "开始抢购",
				OnClicked: func() {
					//comm  := make(chan string, 11)
					db.Submit()
					//抢购
					p := 9515
					for i := 0; i < 5; i++ {
						go selenium.SnapUp(snapUp.SnapUpPageUrl, snapUp.UserName, snapUp.Password, p-1, nil)
					}
					select {}
					////输出日志
					//go func() {
					//	for t := range comm{
					//		outLog = t +outLog
					//		outTE.SetText(outLog)
					//	}
					//}()
				},
			},
			declarative.TextEdit{
				AssignTo: &outTE,
				ReadOnly: true,
				VScroll:  true,
			},
		},
	}.Run()); err != nil {
		log.Fatal(err)
	}
}
