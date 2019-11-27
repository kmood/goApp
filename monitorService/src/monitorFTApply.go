package main

import (
	"fmt"
	"github.com/hprose/hprose-golang/rpc"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/firefox"
	"os"
	"strings"
	"time"
)

func MonitorStatus(port int) {
	const (
		// These paths will be different on your system.
		seleniumPath    = "C:\\selenium-server-standalone-3.4.jar"
		geckoDriverPath = "C:\\geckodriver.exe"
	)

	opts := []selenium.ServiceOption{
		//selenium.StartFrameBuffer(),           // Start an X frame buffer for the browser to run in.
		selenium.GeckoDriver(geckoDriverPath), // Specify the path to GeckoDriver in order to use Firefox.
		selenium.Output(os.Stderr),            // Output debug information to STDERR.

	}
	//selenium.SetDebug(true)

	service, err := selenium.NewSeleniumService(seleniumPath, port, opts...)
	if err != nil {
		panic(err) // panic is used only as an example and is not otherwise recommended.
	}
	defer service.Stop()

	// Connect to the WebDriver instance running locally.
	caps := selenium.Capabilities{
		"browserName": "firefox",
	}
	//禁止图片加载，加快渲染速度
	imagCaps := map[string]interface{}{
		"profile.managed_default_content_settings.images": 2,
	}
	firefoxCaps := firefox.Capabilities{
		Prefs: imagCaps,
		Args: []string{
			//"--headless", // 设置Chrome无头模式，在linux下运行，需要设置这个参数，否则会报错
			//"--no-sandbox",
			"--user-agent=Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.36", // 模拟user-agent，防反爬
		},
	}
	caps.AddFirefox(firefoxCaps)
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://172.16.100.99:%d/wd/hub", port))
	if err != nil {
		panic(err)
	}
	defer wd.Quit()
	f := false
	pageAddress := ""
	for true {
		f, pageAddress = checkPage(wd, "1", "10")
		if f {
			break
		}
		f, pageAddress = checkPage(wd, "2", "10")
		if f {
			break
		}
		f, pageAddress = checkPage(wd, "3", "10")
		if f {
			break
		}
		time.Sleep(time.Second * 10)
	}
	//发送短信
	SendMsg("17611104164", "监测到E级教练员报名开始，请登录以下网址报名："+pageAddress)
}

func SendMsg(phone, content string) {
	client := rpc.NewHTTPClient("http://services.geo-compass.com/messageCat/HproseServer")
	var helloService *HelloService
	client.UseService(&helloService)
	fmt.Println(helloService.SendSingleSMS("test", "test", phone, content))
}

type HelloService struct {
	SendSingleSMS func(string, string, string, string) (string, error)
}

func checkPage(wd selenium.WebDriver, page string, pageSize string) (bool, string) {
	ft := false
	s := "http://010.61hd.net/train/index?page=" + page + "&per-page=" + pageSize
	if err := wd.Get(s); err != nil {
		panic(err)
	}
	// Get a reference to the text box containing code.
	tableEle, err := wd.FindElement(selenium.ByClassName, "table_set")
	elems, err := tableEle.FindElements(selenium.ByTagName, "a")
	if err != nil {
		panic(err)
	}
	for _, elem := range elems {
		s, err := elem.Text()
		if err != nil {
			panic(err)
		}
		s_ := "申请报名"
		//s_ := "未开始报名"
		if strings.Contains(s, s_) {
			ft = true
			break
		}
	}
	return ft, s
}
