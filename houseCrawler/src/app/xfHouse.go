package app

import "time"

type XFHouseInfo struct {
	ID int
	Title string //标题
	DetailURI  string//详情uri
	CrawleURL string//主站url
	CrawlerTime time.Time //爬取时间
	Cjsj       string
	HouseInfo
}

func NewXFHouseInfo() *XFHouseInfo {
	return &XFHouseInfo{}
}
