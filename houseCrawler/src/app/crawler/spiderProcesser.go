package crawler

type SpiderProcesser interface {
	//1.Process 解析响应的字节码
	Process(responseBytes []byte) error
	//2.数据处理器
	DataHandler() error
	//3.数据入库
	DataStore() error

}
