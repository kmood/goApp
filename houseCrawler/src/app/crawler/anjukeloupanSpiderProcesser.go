package crawler

import "app"

type AnjukeProcesser struct {
	AnjukeHouseInfo []*CrawlerLoupanInfo
	HouseInfo []* app.XFHouseInfo
	Url string
}

func (*AnjukeProcesser) Process(responseBytes []byte) error {
	panic("implement me")
}

func (*AnjukeProcesser) DataHandler() error {
	panic("implement me")
}

func (*AnjukeProcesser) DataStore() error {
	panic("implement me")
}

func NewAnjukeProcesser() *AnjukeProcesser {
	return &AnjukeProcesser{}
}

