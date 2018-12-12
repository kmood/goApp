package main

import (
	"fmt"
	"utils"
)

const xLIANJIA_CITY_URL = "https://bj.lianjia.com/city/"
const TEST_FILE_PATH = "C:\\Users\\admin\\Desktop\\crawler\\lianjia_city.html"

func main() {

	//node := utils.ReadNodeByHttp(xLIANJIA_CITY_URL)
	//fmt.Println(node.FirstChild.NextSibling)
	//bodyNode := utils.GetBodyNode(utils.GetHeadHtml(node))

	//获取城市url数据结构
	cityUrlList := &utils.CityUrlList{"https://fs.lianjia.com/city/", make([]utils.CityUrl, 0)}
	cityUrlList.HtmlParse()
	fmt.Println(cityUrlList)

	//file := utils.ReadNodeByFile(TEST_FILE_PATH)

	//fmt.Println(file.FirstChild)

	//fmt.Println(utils.FindNodeByAtrr("class","city_selection_section", utils.GetBodyNode(file)))

}
