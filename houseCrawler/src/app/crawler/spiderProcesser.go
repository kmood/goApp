package crawler

type SpiderProcesser interface {
	//1.Process 解析响应的字节码
	Process(responseBytes []byte) error
	//2.数据处理器 直接入库后期做数据清理
	DataHandler() error
	//3.数据入库
	DataStore() error

}
