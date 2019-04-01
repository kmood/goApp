package selenium

import (
	"fmt"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	"time"
)

const (
	seleniumPath = `D:\Program Files\chormdriver\chromedriver.exe`
	//port         = 9515
	TimeFormat   = "2006-01-02 15:04:05"
	USER_NAME    = "17611104164"
	PASS_WORD    = "BaiChao0305"
	HW_LOGIN_URL = "https://hwid1.vmall.com/CAS/portal/login.html?validated=true&themeName=red&service=https%3A%2F%2Fwww.vmall.com%2Faccount%2Facaslogin%3Furl%3Dhttps%253A%252F%252Fwww.vmall.com%252F&loginChannel=26000000&reqClientType=26&lang=zh-cn"
	//PRODUCT_PAGE_URL = "https://www.vmall.com/product/10086785341226.html#1008643150834211"//正常下单
	PRODUCT_PAGE_URL = "https://www.vmall.com/product/10086134839130.html#1008669544240111" //抢购
)

func SubmitOrder(webPage selenium.WebDriver, comm chan string) error {

	//comm <- time.Now().Format(TimeFormat)+"-----------排队中，检测是否可提交订单...---------------\r\n"
	//*[@id='checkoutSubmit']  正常提交订单
	//submitButton, e := webPage.FindElement(selenium.ByXPATH, "//*[@id='checkoutSubmit']")
	//for submitButton == nil{
	//	submitButton, e = webPage.FindElement(selenium.ByXPATH, "//*[@id='checkoutSubmit']")
	//}
	//抢购
	submitButton, e := webPage.FindElement(selenium.ByXPATH, "//*[@id='submit_order_button']")
	for submitButton == nil {
		submitButton, e = webPage.FindElement(selenium.ByXPATH, "//*[@id='submit_order_button']")
	}
	if e != nil {
		//return errors.New(fmt.Sprintf("获取提交订单button失败，失败原因:%v\r\n",e))
		return e
	}
	webPage.KeyDown(selenium.EndKey)
	webPage.KeyUp(selenium.EndKey)
	e = submitButton.Click()
	if e != nil {
		return e
	}
	time.Sleep(5 * time.Second)
	//comm <- time.Now().Format(TimeFormat)+"-----------提交订单成功，请到登录支付.---------------\r\n"
	return nil
}
func Buy(webPage selenium.WebDriver, snapPageUrl string, comm chan string) error {
	//comm <- time.Now().Format(TimeFormat)+"-----------检测是否可进行抢购...---------------\r\n"
	//打开一个网页
	err := webPage.Get(snapPageUrl)
	if err != nil {
		//fmt.Println("get page faild", err.Error())
		return err
	}
	//检测是否可抢购,策略：立即申购禁止标签为nil时，即可进行申购
	//*[@id='pro-operation']/a
	countDownChildEle, err := webPage.FindElement(selenium.ByXPATH, "//*[@id='pro-operation']/a[@class='product-button02 disabled']")
	for countDownChildEle != nil { //出现抢购按钮跳出循环
		countDownChildEle, _ = webPage.FindElement(selenium.ByXPATH, "//*[@id='pro-operation']/a[@class='product-button02 disabled']")
	}

	//*[@id='pro-operation']/a[2]  正常页面下单
	//buyButton, err := webPage.FindElement(selenium.ByXPATH,"//*[@id='pro-operation']/a[2]")

	//*[@id="pro-operation"]/a 抢购
	buyButton, err := webPage.FindElement(selenium.ByXPATH, "//*[@id='pro-operation']/a")
	for buyButton == nil {
		buyButton, err = webPage.FindElement(selenium.ByXPATH, "//*[@id='pro-operation']/a")
	}
	if err != nil {
		//fmt.Println("获取属性失败", err)
		return err
	}
	err = buyButton.Click()
	if err != nil {
		//fmt.Println("点击立即购买按钮失败",err)
		return err
	}
	//comm <- time.Now().Format(TimeFormat)+"-----------抢购成功，排队中...---------------\r\n"
	return nil
}

func Login(webPage selenium.WebDriver, userName, passWord string, comm chan string) error {
	//comm <- time.Now().Format(TimeFormat)+"-----------开始登录---------------\r\n"
	//打开网页
	error := webPage.Get(HW_LOGIN_URL)
	if error != nil {
		//fmt.Sprintf("打开登录页报错，错误信息为：%v",error)
		return error
	}
	userNameInput, error := webPage.FindElement(selenium.ByXPATH, "//*[@id='login_userName']")
	for userNameInput == nil {
		userNameInput, error = webPage.FindElement(selenium.ByXPATH, "//*[@id='login_userName']")
	}
	if error != nil {
		//fmt.Sprintf("获取登录input属性失败，错误信息为：%v",error)
		return error
	}
	userNameInput.Click()
	error = userNameInput.SendKeys(userName)
	if error != nil {
		//fmt.Sprintf("输入用户名失败，错误信息为：%v",error)
		return error
	}

	////*[@id="login_password"]
	passWordInput, error := webPage.FindElement(selenium.ByXPATH, "//*[@id='login_password']")
	for passWordInput == nil {
		passWordInput, error = webPage.FindElement(selenium.ByXPATH, "//*[@id='login_password']")
	}
	if error != nil {
		//fmt.Sprintf("获取登录input属性失败，错误信息为：%v",error)
		return error
	}
	passWordInput.Click()
	error = passWordInput.SendKeys(passWord)
	if error != nil {
		//fmt.Sprintf("输入密码失败，错误信息为：%v",error)
		return error
	}
	//登录按钮
	lb, error := webPage.FindElement(selenium.ByXPATH, "//*[@id='btnLogin']")
	for lb == nil {
		lb, error = webPage.FindElement(selenium.ByXPATH, "//*[@id='btnLogin']")
	}
	if error != nil {
		//fmt.Sprintf("获取登录button失败，错误信息为：%v",error)
		return error
	}
	lb.Click()

	//comm <- time.Now().Format(TimeFormat)+"-----------登录成功---------------\r\n"
	return nil
}
func SnapUp(snapPageUrl, userName, passWord string, port int, comm chan string) {
	//defer close(comm)
	//如果seleniumServer没有启动，就启动一个seleniumServer所需要的参数，可以为空，示例请参见https://github.com/tebeka/selenium/blob/master/example_test.go
	opts := []selenium.ServiceOption{}
	//opts := []selenium.ServiceOption{
	//    selenium.StartFrameBuffer(),           // Start an X frame buffer for the browser to run in.
	//    selenium.GeckoDriver(geckoDriverPath), // Specify the path to GeckoDriver in order to use Firefox.
	//}

	//selenium.SetDebug(true)
	service, err := selenium.NewChromeDriverService(seleniumPath, port, opts...)
	if nil != err {
		fmt.Println("start a chromedriver service falid", err.Error())
		return
	}
	//注意这里，server关闭之后，chrome窗口也会关闭
	defer service.Stop()

	//链接本地的浏览器 chrome
	caps := selenium.Capabilities{
		"browserName": "chrome",
	}

	//禁止图片加载，加快渲染速度
	imagCaps := map[string]interface{}{
		"profile.managed_default_content_settings.images": 2,
	}
	chromeCaps := chrome.Capabilities{
		Prefs: imagCaps,
		Path:  "",
		Args: []string{
			//"--headless", // 设置Chrome无头模式，在linux下运行，需要设置这个参数，否则会报错
			//"--no-sandbox",
			"--user-agent=Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.36", // 模拟user-agent，防反爬
		},
	}
	//以上是设置浏览器参数
	caps.AddChrome(chromeCaps)

	// 重新调起chrome浏览器
	webPage, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		//comm <- fmt.Sprintf("-------------------------connect to the webDriver faild%v\r\n", err.Error())
		return
	}
	defer webPage.Close()
	//登录
	err = Login(webPage, userName, passWord, comm)
	if err != nil {
		//comm <- fmt.Sprintf("-------------------------login error !!!!!!!!!%v\r\n", err.Error())
		return
	}
	//for i:=0;i<20 ;i++  {
	//go func() {
	err = Buy(webPage, snapPageUrl, comm)
	if err != nil {
		//comm <- fmt.Sprintf("-------------------------Buy error !!!!!!!!!%v\r\n", err.Error())
		return
	}

	err = SubmitOrder(webPage, comm)
	if err != nil {
		//comm <- fmt.Sprintf("-------------------------login error !!!!!!!!!%v\r\n", err.Error())
		return
	}
	//}()
	//}
	return
}
