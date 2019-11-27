package main

import (
	"flag"
	"os"
	"time"

	"github.com/DazzlingSun/monitorService/src/project/conf"
	"github.com/DazzlingSun/monitorService/src/project/http"
	"github.com/DazzlingSun/monitorService/src/project/service"
	"github.com/DazzlingSun/monitorService/src/basic/log"
	"github.com/DazzlingSun/monitorService/src/basic/net/trace"
	"github.com/DazzlingSun/monitorService/src/basic/os/signal"
	"github.com/DazzlingSun/monitorService/src/basic/syscall"
)

var (
	s *service.Service
)

func main() {
	flag.Parse()
	if err := conf.Init(); err != nil {
		log.Error("conf.Init() error(%v)", err)
		panic(err)
	}
	log.Init(conf.Conf.Log)
	defer log.Close()
	trace.Init(conf.Conf.Tracer)
	defer trace.Close()
	s = service.New(conf.Conf)
	http.Init(conf.Conf, s)
	log.Info("push-admin start")
	signalHandler()
}

func signalHandler() {
	var (
		ch = make(chan os.Signal, 1)
	)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		si := <-ch
		switch si {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGSTOP, syscall.SIGINT:
			time.Sleep(time.Second * 2)
			log.Info("get a signal %s, stop the push-admin process", si.String())
			s.Close()
			s.Wait()
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
