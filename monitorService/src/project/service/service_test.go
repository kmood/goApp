package service

import (
	"flag"
	"path/filepath"
	"testing"
	"time"

	"github.com/DazzlingSun/monitorService/src/project/conf"

	. "github.com/smartystreets/goconvey/convey"
)

var svr *Service

func init() {
	dir, _ := filepath.Abs("../cmd/activity-admin-test.toml")
	flag.Set("conf", dir)
	conf.Init()
	svr = New(conf.Conf)
	time.Sleep(time.Second)
}

func WithService(f func(s *Service)) func() {
	return func() {
		Reset(func() {})
		f(svr)
	}
}

func Test_Service(t *testing.T) {
	Convey("service test", t, WithService(func(s *Service) {
		s.Wait()
		s.Close()
	}))
}
