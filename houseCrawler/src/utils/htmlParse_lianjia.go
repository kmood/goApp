package utils

import (
	"fmt"
	"golang.org/x/net/html"
	"strconv"
	"strings"
)

const CITY_NODE_CLASSNAME = "city_list_ul"

type CityUrlList struct {
	LianjiaCityURL string
	CityUrls       []CityUrl
}
type CityUrl struct {
	Province   string
	CityUrlMap map[string]string
}

func (cityurlList *CityUrlList) GetCityUrlByName(province, cityName string) string {
	urls := cityurlList.CityUrls
	for i := 0; i < len(urls); i++ {
		cityUrl := urls[i]
		if cityUrl.Province == province {
			return cityUrl.CityUrlMap[cityName]
		}
	}
	return ""
}
func visit(datamap map[string]string, node *html.Node, attrName, attrValue string) {

	if node == nil {
		return
	}
	if node.Type == html.ElementNode && node.Data == "a" {
		if len(attrValue) != 0 && len(attrName) != 0 {
			for _, a := range node.Attr {
				if a.Key == attrName && a.Val == attrValue {
					for _, a := range node.Attr {
						if a.Key == "href" {
							datamap[node.FirstChild.Data] = a.Val
						}
					}
				}
			}
		} else {
			for _, a := range node.Attr {
				if a.Key == "href" {
					datamap[node.FirstChild.Data] = a.Val
				}
			}
		}
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		visit(datamap, c, "", "")
	}
}
func (cityurlList *CityUrlList) HtmlParse() {
	cityDoc := ReadNodeByHttp(cityurlList.LianjiaCityURL)
	BodyNode := GetBodyNode(cityDoc)
	//fmt.Println(BodyNode)
	cityListNode := FindNodeByAtrr("class", CITY_NODE_CLASSNAME, BodyNode)
	node := cityListNode.FirstChild
	//
	for cityNode := node; cityNode != nil; cityNode = cityNode.NextSibling {
		if cityNode.Type != html.ElementNode {
			continue
		}
		var cityurl CityUrl
		provinceNode := FindNodeByAtrr("class", "city_list_tit c_b", cityNode)
		cityurl.Province = provinceNode.FirstChild.Data
		cityUrlMap := make(map[string]string)
		visit(cityUrlMap, cityNode, "", "")
		cityurl.CityUrlMap = cityUrlMap
		cityurlList.CityUrls = append(cityurlList.CityUrls, cityurl)
	}
}

type RentHouseInfo struct {
	CityUrl   string
	City      string
	Province  string
	HouseInfo []HouseInfo
}

func (rentHouseInfo *RentHouseInfo) HtmlParse() {
	cityDoc := ReadNodeByHttp(rentHouseInfo.CityUrl + "zufang/")
	bodyNode := GetBodyNode(cityDoc)
	var totalPage int
	//获取内容主体节点
	contentNode := FindNodeByAtrr("class", "content__article", bodyNode)
	zfTotalNode := FindNodeByAtrr("class", "content__title--hl", contentNode)
	zfTotalNum, er := strconv.ParseInt(zfTotalNode.FirstChild.Data, 0, 32)
	if er != nil {
		fmt.Printf("解析租房总数报错:%v", er)
	}
	if zfTotalNum%30 == 0 { //计算总页数
		totalPage = int(zfTotalNum / 30)
	} else {
		totalPage = int(zfTotalNum/30) + 1
	}
	for i := 1; i < totalPage; i++ {
		nodeByHttp := ReadNodeByHttp(fmt.Sprintf(rentHouseInfo.CityUrl+"zufang/pg%d/", i))
		bodyNode = GetBodyNode(nodeByHttp)
		contentListNode := FindNodeByAtrr("class", "content__list", bodyNode)

		for node := contentListNode.FirstChild; node != nil; node = node.NextSibling {
			houseInfo := NewHouseInfo()
			if node.Type == html.ElementNode {
				//获取详情页url
				node1 := FindNodeByAtrr("class", "content__list--item--aside", node)
				for _, attr := range node1.Attr {
					if attr.Key == "href" {
						houseInfo.DetailPageUrl = rentHouseInfo.CityUrl + attr.Val
					}
				}
				//获取标题、图片路径  !!!!!此处class 的值与页面看到的不一致  页面为 " lazyloaded"
				node2 := FindNodeByAtrr("class", "lazyload", node1)
				for _, attr := range node2.Attr {
					if attr.Key == "alt" {
						houseInfo.Title = attr.Val
					}
					if attr.Key == "src" {
						houseInfo.PictureUrl = attr.Val
					}
				}
				//获取 地址
				node3 := FindNodeByAtrr("class", "content__list--item--des", node)
				i := 0
				for firstChild := node3.FirstChild; firstChild != nil; firstChild = firstChild.NextSibling {
					if firstChild.Data == "a" {
						houseInfo.Address += firstChild.FirstChild.Data
					}
					//获取房屋尺寸，朝向，房屋类型
					if firstChild.Type == html.TextNode {
						switch i {
						case 1:
							replace := strings.Replace(firstChild.Data, " ", "", -1)
							replace = strings.Replace(replace, "\n", "", -1)
							houseInfo.Size = replace
							break
						case 2:
							replace := strings.Replace(firstChild.Data, " ", "", -1)
							replace = strings.Replace(replace, "\n", "", -1)
							houseInfo.Direction = replace
							break
						case 3:
							replace := strings.Replace(firstChild.Data, " ", "", -1)
							replace = strings.Replace(replace, "\n", "", -1)
							houseInfo.HouseType = replace
							break
						}
					}
					if firstChild.Data == "i" {
						i++
					}
				}
				//获取房屋价格
				node4 := FindNodeByAtrr("class", "content__list--item-price", node)
				houseInfo.Price = strings.Replace(node4.FirstChild.FirstChild.Data, " ", "", -1)
				houseInfo.PriceUnit = strings.Replace(node4.FirstChild.NextSibling.Data, " ", "", -1)
				//获取发布时间
				node5 := FindNodeByAtrr("class", "content__list--item--time oneline", node)
				houseInfo.PublishTime = node5.FirstChild.Data
				//设置基本的信息
				houseInfo.CityUrl = rentHouseInfo.CityUrl
				rentHouseInfo.HouseInfo = append(rentHouseInfo.HouseInfo, *houseInfo)
			}
		}
	}
}

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

func NewHouseInfo() *HouseInfo {
	return &HouseInfo{}
}
func NewRentHouseInfo() *RentHouseInfo {
	return &RentHouseInfo{}
}
func (r *RentHouseInfo) Insert() (int64, error) {
	db, e := GetDB()
	defer db.Close()
	tx, e := db.Begin()
	if e != nil {
		return 0, e
	}
	i := 0
	var insertRow int64

	sql := "insert into houseinfo(title,address,price,priceunit,housetype,size,rentway,position,publishtime,direction,pictureurl,cityurl,detailpageurl)  values"
	for _, houseInfo := range r.HouseInfo {
		sql += " ( '" + houseInfo.Title + "','" + houseInfo.Address + "','" + houseInfo.Price + "','" + houseInfo.PriceUnit + "','" + houseInfo.HouseType + "','" + houseInfo.Size + "','" + houseInfo.RentWay + "'," +
			"'" + houseInfo.Position + "','" + houseInfo.PublishTime + "','" + houseInfo.Direction + "','" + houseInfo.PictureUrl + "','" + houseInfo.CityUrl + "','" + houseInfo.DetailPageUrl + "'),"
		i++
		if i == 5000 {
			i = 0
			sql = strings.TrimSuffix(sql, ",")
			result, e1 := db.Exec(sql)
			//事务提交
			tx.Commit()
			if e1 != nil {
				return 0, e1
			}
			i, e2 := result.RowsAffected()
			if e2 != nil {
				return insertRow, e2
			}
			sql = "insert into houseinfo(title,address,price,priceunit,housetype,size,rentway,position,publishtime,direction,pictureurl,cityurl,detailpageurl)  values"
			insertRow += i
		}
	}
	if len(r.HouseInfo)%5000 != 0 {
		sql = strings.TrimSuffix(sql, ",")
		result, e1 := db.Exec(sql)
		//事务提交
		tx.Commit()
		if e1 != nil {
			return 0, e1
		}
		i, e2 := result.RowsAffected()
		if e2 != nil {
			return insertRow, e2
		}
		insertRow += i
	}
	return insertRow, nil
}
