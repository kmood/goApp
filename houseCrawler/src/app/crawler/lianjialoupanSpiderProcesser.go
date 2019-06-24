package crawler

import (
	"app"
	"app/dao"
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/jinzhu/gorm"
	"strconv"
	"strings"
	"time"
	"utils"
)

// LianjiaLoupanPageProcesser 链家楼盘页面解析器
type LianjiaProcesser struct {
	LianjiaHouseInfo []*CrawlerLoupanInfo
	HouseInfo []* app.XFHouseInfo
	Url string

}

func NewLianjiaProcesser() *LianjiaProcesser {
	return &LianjiaProcesser{}
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

		infoPrice := NewCrawlerLianjiaLoupanInfoPrice
		houseInfo.PriceInfo = infoPrice()
		//获取价格
		priceSelect := selectorPrefix + " div > div.resblock-price > div.main-price > span.number"
		price := mainNodeSelector.Find(priceSelect).Text()
		atoi, e := strconv.Atoi(price)
		if e != nil {
			houseInfo.PriceInfo.Price = 0
		}else {
			houseInfo.PriceInfo.Price = atoi
		}
		//价格单位
		houseInfo.PriceInfo.PriceUnit = mainNodeSelector.Find(selectorPrefix +" div > div.resblock-price > div.main-price > span.desc").Text()
		//总价单位
		houseInfo.PriceInfo.TotalPrice = mainNodeSelector.Find(selectorPrefix +" div > div.resblock-price > div.second").Text()
		//建筑面积
		houseInfo.CoveredArea = mainNodeSelector.Find(selectorPrefix + "div > div.resblock-area > span").Text()
		//特点
		houseInfo.Feature = utils.DelInvisibleChar(strings.ReplaceAll(mainNodeSelector.Find(selectorPrefix+" div > div.resblock-tag").Text(), "\n", "/"))
		//创建时间
		houseInfo.PriceInfo.Cjsj = time.Now().Format("2006-01-02 15:04:10")
		//详情页uri
		houseInfo.DetailURI, _ = mainNodeSelector.Find(selectorPrefix + " div > div.resblock-name > a").Attr("href")
		//房屋唯一id
		houseInfo.PriceInfo.HouseID = "lianjia_loupan"+ houseInfo.DetailURI
		houseInfo.HouseID = "lianjia_loupan"+ houseInfo.DetailURI
		//添加
		h.LianjiaHouseInfo = append(h.LianjiaHouseInfo, houseInfo)
		//打印
		//println(fmt.Sprintf("%+v", houseInfo))
		i++
	}
	return nil
}

func (h *LianjiaProcesser) DataHandler() error{
	return nil
}

func  (h *LianjiaProcesser)DataStore()error{
	d := dao.New()
	gorm.AddNamingStrategy(&gorm.NamingStrategy{
		Column: func(name string) string {
			return strings.ToLower(name)
		},
	})
	db := d.ORM
	db.AutoMigrate(&CrawlerLoupanInfoPrice{})
	db.AutoMigrate(&CrawlerLoupanInfo{})
	for _,lhi := range h.LianjiaHouseInfo {
		priceInfo := lhi.PriceInfo
		db.Create(&priceInfo)
		if !db.NewRecord(lhi){
			db.Create(lhi)
		}
	}
	return nil
}
func (h *LianjiaProcesser)Spider(crawleURL string) error{
	newCollector := colly.NewCollector()
	newCollector.OnResponse(func(response *colly.Response) {
		document, e := goquery.NewDocumentFromReader(bytes.NewReader(response.Body))
		if e != nil {
			return
		}
		houseNumStr := document.Find("body > div.resblock-list-container.clearfix > div.resblock-have-find > span.value").Text()
		houseNum, e := strconv.Atoi(houseNumStr)
		if e != nil {
			return
		}
		h.Url = crawleURL
		//天津链家 https://tj.fang.lianjia.com/loupan/pg3/
		if houseNum % 10 !=0 {
			houseNum += 10
		}
		for i := 1; i < houseNum/10; i++ {
			collector := colly.NewCollector()
			url := crawleURL + "/loupan/pg"+strconv.Itoa(i)+"/"
			println(url)
			collector.OnResponse(func(response *colly.Response) {
				//解析页面
				h.Process(response.Body)
			})
			collector.Visit(url)
		}
		//数据处理
		h.DataHandler()
		//入库
		h.DataStore()
	})
	newCollector.Visit(crawleURL+ "/loupan/pg1/")
	return nil
}


