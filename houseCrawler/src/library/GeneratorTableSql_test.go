package library

import "testing"

type HouseInfo struct {
	Id            string `comment_size:"数据id_50"`
	Title         string `comment_size:"链家标题_50"`
	Address       string `comment_size:"房屋地址_100"`
	Price         string `comment_size:"价格_10"`
	PriceUnit     string `comment_size:"价格单位_10"`
	HouseType     string `comment_size:"房子类型_10"`
	Size          string `comment_size:"房子大小_10"`
	RentWay       string `comment_size:"租赁方式_10"`
	Position      string `comment_size:"详细位置_100"`
	PublishTime   string `comment_size:"发布时间_50"`
	Direction     string `comment_size:"房屋朝向_10"`
	PictureUrl    string `comment_size:"图片路径_100"`
	CityUrl       string `comment_size:"链家城市url_200"`
	DetailPageUrl string `comment_size:"链家当前房屋详情链接_200"`
}
type LianjiaHouseInfo struct {
	HouseID     string `comment_size:"房屋唯一id_100"`//标题
	Title       string `comment_size:"标题_100"`//标题"
	Position    string `comment_size:"实际位置_100"`//实际位置
	DoorModel   string `comment_size:"户型_100"`//户型
	Price       int    `comment_size:"价格_100"`//价格
	PriceUnit   string `comment_size:"价格单位_100"`//价格单位
	TotalPrice  string `comment_size:"总价_100"`//总价
	CoveredArea string `comment_size:"建筑面积_100"`// 建筑面积
	Feature     string `comment_size:"特点_100"`//特点
	Cjsj        string `comment_size:"爬取时间_100"`//爬取时间
	DetailURI   string `comment_size:"查看详情url_100"`//查看详情url
}
func TestGeneratorTableSql(t *testing.T) {
	sql := GeneratorTableSql(LianjiaHouseInfo{},"crawler_lianjia_loupan_info","爬取-链家-楼盘-数据信息")
	t.Log(sql)
}
