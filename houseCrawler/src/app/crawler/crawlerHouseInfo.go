package crawler


type CrawlerLoupanInfo struct {
	HouseID 	string`gorm:"primary_key"`
	PriceInfo  		*CrawlerLoupanInfoPrice
	Title       string `comment_size:标题_100`//标题
	Position    string `comment_size:实际位置_100`//实际位置
	DoorModel   string `comment_size:户型_100`//户型
	CoveredArea string `comment_size:建筑面积_100`// 建筑面积
	Feature     string `comment_size:特点_100`//特点
	DetailURI   string `comment_size:查看详情url_100`//查看详情url
	CrawlerTime string `comment_size:抓取时间_100`
	OnSale 		string `comment_size:在售(0否1是)_1`
}
type CrawlerLoupanInfoPrice struct {
	HouseID     string `comment_size:房屋唯一id_100`//标题
	Price       int    `comment_size:价格_10`//价格
	PriceUnit   string `comment_size:价格单位_100`//价格单位
	RimPrice    int    `comment_size:周边均价_10`
	TotalPrice  string `comment_size:总价_100`//总价
	Cjsj        string `comment_size:爬取时间_100`//爬取时间
}

func NewCrawlerLianjiaLoupanInfoPrice() *CrawlerLoupanInfoPrice {
	return &CrawlerLoupanInfoPrice{}
}

func NewLianjiaHouseInfo() *CrawlerLoupanInfo {
	return &CrawlerLoupanInfo{}
}
