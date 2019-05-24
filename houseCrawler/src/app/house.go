package app

//const (
//	ZF  = iota //租房
//	XF         //现房
//	ESF        //二手房
//)

type HouseInfo struct {
	DoorModel  string // 户型
	Lat        float64
	Lon        float64
	Position   string
	Feature     string //特点
	Price      int
}

func NewHouseInfo() *HouseInfo {
	return &HouseInfo{}
}
