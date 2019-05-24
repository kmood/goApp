package main

const xLIANJIA_CITY_URL = "https://bj.lianjia.com/city/"
const TEST_FILE_PATH = "C:\\Users\\admin\\Desktop\\crawler\\lianjia_city.html"

type MyPageProcesser struct {
}

func main() {
	////抓取数据
	//rentHouseInfo := utils.NewRentHouseInfo()
	//rentHouseInfo.HouseInfo = make([]utils.HouseInfo, 0)
	//rentHouseInfo.CityUrl = "https://bj.lianjia.com/"
	//
	//rentHouseInfo.HtmlParse()
	//i, e := rentHouseInfo.Insert()
	//if e != nil {
	//	fmt.Println("插入到数据库失败", e)
	//}
	//fmt.Printf("-------新建%d条数据----------", i)
	//
	////数据库表生成
	//fmt.Println(utils.GeneratorTableSql(utils.HouseInfo{}))


}

//func getUniqe(arrs []int) int {
//	ret := arrs[0]
//	for _, arr := range arrs {
//		ret = ret ^ arr
//	}
//	return ret
//}
