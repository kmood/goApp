package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/DazzlingSun/monitorService/src/basic/log"
	"github.com/DazzlingSun/monitorService/src/basic/net/trace"
	"github.com/DazzlingSun/monitorService/src/basic/rate/limit/bench/stress/conf"
	"github.com/DazzlingSun/monitorService/src/basic/rate/limit/bench/stress/http"
)

func main() {
	flag.Parse()
	if err := conf.Init(); err != nil {
		log.Error("conf.Init() error(%v)", err)
		panic(err)
	}
	// init log
	log.Init(conf.Conf.Log)
	defer log.Close()
	log.Info("stress start")
	// init trace
	trace.Init(conf.Conf.Tracer)
	defer trace.Close()
	// ecode init
	//	ecode.Init(conf.Conf.Ecode)
	// service init
	http.Init(conf.Conf)
	// init pprof conf.Conf.Perf
	go func() {
		// init signal
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
		for {
			s := <-c
			fmt.Println("go sig!!!!!!!!")

			log.Info("stress get a signal %s", s.String())
			switch s {
			case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
				log.Info("stress exit")
				os.Exit(0)
				return
			case syscall.SIGHUP:
				os.Exit(0)

			default:
				os.Exit(0)

				return
			}
		}
	}()
	ch := make(chan struct{})
	<-ch
}
