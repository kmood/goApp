package crawler

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"strconv"
	"strings"
	"time"
	"utils"
)

// LianjiaLoupanPageProcesser 链家楼盘页面解析器
type LianjiaProcesser struct {
	LianjiaHouseInfo []*LianjiaHouseInfo
	Route            *Route
}
type LianjiaHouseInfo struct {
	Title       string //标题
	Position    string //实际位置
	DoorModel   string //户型
	Price       int    //价格
	CoveredArea string // 建筑面积
	Feature     string //特点
	Cjsj        string //爬取时间
	DetailURI   string //查看详情url
}

func NewLianjiaHouseInfo() *LianjiaHouseInfo {
	route := NewRoute()
	route.Name = "链家"
	route.Url = "https://bj.lianjia.com"
	return &LianjiaHouseInfo{}
}

func NewLianjiaProcesser() *LianjiaProcesser {
	return &LianjiaProcesser{}
}
func (h *LianjiaProcesser) InitRoute() error {
	c := colly.NewCollector()
	// On every a element which has href attribute call callback
	c.OnResponse(func(response *colly.Response) {
		document, e := goquery.NewDocumentFromReader(bytes.NewReader(response.Body))
		if e != nil {
			return
		}
		houseNumStr := document.Find("body > div.resblock-list-container.clearfix > div.resblock-have-find > span.value").Text()
		houseNum, e := strconv.Atoi(houseNumStr)
		if e != nil {
			return
		}
		h.Route.Url = "https://tj.fang.lianjia.com"
		h.Route.Name = "链家"
		//链家 https://tj.fang.lianjia.com/loupan/pg3/
		for i := 1; i < houseNum/10; i++ {
			h.Route.Uris = append(h.Route.Uris, "/loupan/pg"+strconv.Itoa(i)+"/")
		}
	})
	c.Visit("https://tj.fang.lianjia.com/loupan/pg1/")
	return nil
}

// Parse html dom here and record the parse result that we want to crawl.
// Package goquery (http://godoc.org/github.com/PuerkitoBio/goquery) is used to parse html.
func (h *LianjiaProcesser) Process(responseBytes []byte) error {
	document, e := goquery.NewDocumentFromReader(bytes.NewReader(responseBytes))
	if e != nil {
		return e
	}
	//主节点
	mainNodeSelector := document.Find("body > div.resblock-list-container.clearfix > ul.resblock-list-wrapper")
	mainNode := mainNodeSelector.Nodes[0]
	i := 1
	//
	for node := mainNode.FirstChild; node != nil; node = node.NextSibling {
		houseInfo := NewLianjiaHouseInfo()
		if node.Type != 3 { //非元素标签时跳过
			continue
		}
		selectorPrefix := "li:nth-child(" + fmt.Sprintf("%d", i) + ") > "
		//find  查询到的节点 text（） 去除子节点的所有内容
		//标题
		houseInfo.Title = mainNodeSelector.Find(selectorPrefix + " div > div.resblock-name > a").Text()
		//实际位置  城市
		//body > div.main-nav-container > div > div.main-left-wrapper > a.s-city
		//body > div.main-nav-container > div > div.main-left-wrapper > a.s-city
		cs := document.Find("a[class=\"s-city\"]").Text() //城市
		houseInfo.Position += cs + "-"
		houseInfo.Position += mainNodeSelector.Find(selectorPrefix+" div > div.resblock-location > span:nth-child(1)").Text() + "-"
		houseInfo.Position += mainNodeSelector.Find(selectorPrefix+" div > div.resblock-location > span:nth-child(3)").Text() + "-"
		houseInfo.Position += mainNodeSelector.Find(selectorPrefix + " div > div.resblock-location > a").Text()
		//获取户型
		houseInfo.DoorModel = utils.DelInvisibleChar(mainNodeSelector.Find(selectorPrefix + "div > a").Text())
		//获取价格
		priceSelect := selectorPrefix + " div > div.resblock-price > div.main-price > span.number"
		price := mainNodeSelector.Find(priceSelect).Text()
		atoi, e := strconv.Atoi(price)
		if e != nil {
			fmt.Printf("解析字符串到整型报错", e)
		}
		houseInfo.Price = atoi
		//建筑面积
		houseInfo.CoveredArea = mainNodeSelector.Find(selectorPrefix + "div > div.resblock-area > span").Text()
		//特点
		houseInfo.Feature = utils.DelInvisibleChar(strings.ReplaceAll(mainNodeSelector.Find(selectorPrefix+" div > div.resblock-tag").Text(), "\n", "/"))
		//创建时间
		houseInfo.Cjsj = time.Now().Format("2006-01-02 15:04:10")
		//详情页uri
		houseInfo.DetailURI, _ = mainNodeSelector.Find(selectorPrefix + " div > div.resblock-name > a").Attr("href")
		//添加
		h.LianjiaHouseInfo = append(h.LianjiaHouseInfo, houseInfo)
		//打印
		println(fmt.Sprintf("%+v", houseInfo))
		i++
	}
	return nil
}

func (h *LianjiaProcesser) DataHandler() error {
	return nil
}
