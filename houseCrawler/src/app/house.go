package app

const (
	ZF  = iota //租房
	XF         //现房
	ESF        //二手房
)

type HouseInfo struct {
	Id         string
	DoorModel  string
	Lat        float64
	Lon        float64
	Position   string
	Price      int
	Cjsj       string
	CrawleSite string
	DetailURI  string
	HouseType  string
}

func NewHouseInfo() *HouseInfo {
	return &HouseInfo{}
}
