package crawler

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestLianjiaProcesser_DataStore(t *testing.T) {
	processer := NewLianjiaProcesser()
	price1 := NewCrawlerLianjiaLoupanInfoPrice()
	info1 := NewLianjiaHouseInfo()
	info1.HouseID = "lianjia_loupan_test1"
	price1.HouseID = "lianjia_loupan_test1"
	info1.PriceInfo = price1
	price2 := NewCrawlerLianjiaLoupanInfoPrice()
	info2 := NewLianjiaHouseInfo()
	info2.HouseID = "lianjia_loupan_test2"
	price2.HouseID = "lianjia_loupan_test2"
	info2.PriceInfo = price2
	processer.LianjiaHouseInfo = append(processer.LianjiaHouseInfo, info1)
	processer.LianjiaHouseInfo = append(processer.LianjiaHouseInfo, info2)
	processer.DataStore()

}
func TestLianjiaProcesser_Process(t *testing.T) {

}

func TestLianjiaProcesser_Spider(t *testing.T) {
	processer := NewLianjiaProcesser()
	e := processer.Spider("https://tj.fang.lianjia.com/")
	if e != nil {
		t.Error(e)
	}
	logrus.Printf("数据%v",fmt.Sprintf("%v",processer.LianjiaHouseInfo))
}
