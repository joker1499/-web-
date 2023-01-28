package liner

import (
	"bufio"
	"fmt"
	"github.com/antchfx/htmlquery"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io"
	"math/rand"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

const (
	//指定对应浏览器的driverD:\360驱动大师目录
	chromeDriverPath = "D:\\360驱动大师目录\\chromedriver.exe"
	port             = 8080
)

type Msg struct {
	Id  int
	Msg string
}

func getRandstring(length int) string {
	if length < 1 {
		return ""
	}
	char := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	charArr := strings.Split(char, "")
	charlen := len(charArr)
	ran := rand.New(rand.NewSource(time.Now().Unix()))
	var rchar string = ""
	for i := 1; i <= length; i++ {
		rchar = rchar + charArr[ran.Intn(charlen)]
	}
	return rchar
}

// 随机文件名
func RandFileName(fileName string) string {
	randStr := getRandstring(16)
	return randStr + filepath.Ext(fileName)
}

func Fire(url string, params string) {

	source := start(url)
	file := "测试.txt"
	name := RandFileName(file)
	tt, _ := os.Create("temp/" + name + "01")
	defer tt.Close()

	mm, _ := os.Create("temp/" + name + "02")
	defer mm.Close()

	node, err := htmlquery.Parse(strings.NewReader(source))

	if err != nil {
		fmt.Println(err)
	}

	for _, e := range htmlquery.Find(node, "//*[@class=\"p1\"]") {
		val := htmlquery.InnerText(e)
		val = strings.Replace(val, "\n", "", -1)
		val = strings.TrimSpace(val)
		tt.Write([]byte(val))
	}

	for _, e := range htmlquery.Find(node, "//*[@id=\"a_bibliography\"]/p") {
		val := htmlquery.InnerText(e)
		val = strings.Replace(val, "\n", "", -1)
		val = strings.TrimSpace(val)
		mm.Write([]byte(val + "\n"))
	}

	res := query(params, name)

	if len(res) != 0 {
		march(res, name)
	}

}

func query(s string, f string) string {

	file, _ := os.Open("temp/" + f + "02")
	defer file.Close()
	r := bufio.NewReader(file)
	for {
		data2, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println("read err", err.Error())
			break
		}

		fmt.Println(strings.Contains(data2, s))
		if strings.Contains(data2, s) {
			compile, _ := regexp.Compile("\\[\\d+]")
			allString := compile.FindAllString(data2, -1)
			//fmt.Println(allString)
			newStr := strings.TrimLeft(allString[0], "[")
			newStr = strings.TrimRight(newStr, "]")
			//fmt.Println(newStr)
			return newStr
		}
	}
	return ""
}

func march(s string, f string) {

	dsn := "root:root@tcp(124.70.14.47:3306)/alba?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	fmt.Println("db = ", db)
	fmt.Println("err = ", err)

	file, _ := os.Open("temp/" + f + "01")
	defer file.Close()
	r := bufio.NewReader(file)
	data2, _ := r.ReadString('\n')
	fmt.Println(s)

	compile, _ := regexp.Compile("(.*?)\\[" + s + "\\](.*?)。")
	allString := compile.FindAllString(data2, -1)

	split := strings.Split(allString[0], "。")

	if strings.HasPrefix(split[len(split)-2], "[") {
		fmt.Println(split[len(split)-3])
		msg := Msg{Msg: split[len(split)-3]}
		result := db.Create(&msg)
		fmt.Println(msg.Id)       // 返回插入数据的主键
		fmt.Println(result.Error) // 返回 error
		// 返回插入记录的条数
		fmt.Println(result.RowsAffected)
		return
	}
	right := strings.Replace(split[len(split)-2], "["+s+"]", "", -1)
	msg := Msg{Msg: right}
	result := db.Create(&msg)
	fmt.Println(msg.Id)       // 返回插入数据的主键
	fmt.Println(result.Error) // 返回 error
	// 返回插入记录的条数
	fmt.Println(result.RowsAffected)

}

func start(url string) string {
	var opts []selenium.ServiceOption
	//selenium.Output(os.Stderr), // Output debug information to STDERR.
	//SetDebug 设置调试模式
	//selenium.SetDebug(false)
	//在后台启动一个ChromeDriver实例
	service, err := selenium.NewChromeDriverService(chromeDriverPath, port, opts...)
	if err != nil {
		panic(err) // panic is used only as an example and is not otherwise recommended.
	}
	defer service.Stop()

	//连接到本地运行的 WebDriver 实例
	//这里的map键值只能为browserName，源码里需要获取这个键的值，来确定连接的是哪个浏览器
	caps := selenium.Capabilities{"browserName": "chrome"}

	imgCaps := map[string]interface{}{
		"profile.managed_default_content_settings.images": 2,
	}
	chromeCaps := chrome.Capabilities{
		Prefs: imgCaps,
		Path:  "",
		Args: []string{
			//"--headless",
			"--start-maximized",
			//"--window-size=1200x600",
			"--no-sandbox",
			"--user-agent=Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.77 Safari/537.36",
			"--disable-gpu",
			"--disable-impl-side-painting",
			"--disable-gpu-sandbox",
			"--disable-accelerated-2d-canvas",
			"--disable-accelerated-jpeg-decoding",
			"--test-type=ui",
		},
	}
	caps.AddChrome(chromeCaps)
	//NewRemote 创建新的远程客户端，这也将启动一个新会话。 urlPrefix 是 Selenium 服务器的 URL，必须以协议 (http, https, ...) 为前缀。为urlPrefix提供空字符串会导致使用 DefaultURLPrefix,默认访问4444端口，所以最好自定义，避免端口已经被抢占。后面的路由还是照旧DefaultURLPrefix写
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		panic(err)
	}
	defer wd.Quit()

	// 导航到指定网站页面，浏览器默认get方法
	if err := wd.Get(url); err != nil {
		panic(err)
	}

	time.Sleep(3 * time.Second)

	w1, error := wd.FindElement(selenium.ByXPATH, "//*[@id=\"DownLoadParts\"]/div[2]/ul/li[2]/a")
	if error != nil {
		fmt.Println(error)
	}
	time.Sleep(1 * time.Second)
	//fmt.Println(w1)
	w1.Click()
	time.Sleep(5 * time.Second)

	windowHandles, err := wd.WindowHandles()
	if err != nil {
		fmt.Println(err)
	}

	wd.SwitchWindow(windowHandles[1])
	source, err := wd.PageSource()
	if err != nil {
		fmt.Println(err)
	}

	elements, err := wd.FindElements(selenium.ByXPATH, "//*[@class=\"p1\"]")

	if err != nil {
		fmt.Println(err)
	}

	if len(elements) == 0 {
		wd.Refresh()
		time.Sleep(1 * time.Second)
		source, _ = wd.PageSource()
		time.Sleep(1 * time.Second)
	}

	return source
}

func save(db *gorm.DB, msg Msg) {
	result := db.Create(&msg)
	fmt.Println(msg.Id)       // 返回插入数据的主键
	fmt.Println(result.Error) // 返回 error
	// 返回插入记录的条数
	fmt.Println(result.RowsAffected)
}
