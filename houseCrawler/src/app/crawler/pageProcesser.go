package crawler

type PageProcesser interface {
	//InitRoute  解析路由 houseType 使用//ZF = iota 租房	//XF 现房 //ESF 二手房
	InitRoute() error
	//Process 解析响应的字节码
	Process(responseBytes []byte) error
	//数据处理器
	DataHandler() error
}
