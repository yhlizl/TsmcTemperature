package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/spf13/viper"
	"github.com/tebeka/selenium"
)

const (
	chromeDriverPath = "./chromedriver92"
	port             = 8777
)

func click(wd selenium.WebDriver, arr []string) {
	for _, v := range arr {
		fmt.Println("start to click : ", v)
		btn, err := wd.FindElement(selenium.ByCSSSelector, v)
		if err != nil {
			log.Panic(err, v)
		}
		if err := btn.Click(); err != nil {
			log.Panic(err, v)
		}
	}
}

var (
	liststr string
	tempMax float64
	tempMin float64
	emplid  string
)

func init() {
	viper.SetConfigType("toml")
	viper.SetConfigFile("./config/config.toml") // 指定配置文件路徑
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig() // 查找並讀取配置文件
	if err != nil {             // 處理讀取配置文件的錯誤
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	liststr = viper.GetString("list.list")
	tempMax = viper.GetFloat64("temperature.tempMax")
	tempMin = viper.GetFloat64("temperature.tempMin")
	emplid = viper.GetString("profile.emplid")

}
func main() {
	// Start a WebDriver server instance
	opts := []selenium.ServiceOption{
		selenium.Output(os.Stderr), // Output debug information to STDERR.
	}
	selenium.SetDebug(true)
	service, err := selenium.NewChromeDriverService(chromeDriverPath, port, opts...)
	if err != nil {
		panic(err) // panic is used only as an example and is not otherwise recommended.
	}
	defer service.Stop()

	// Connect to the WebDriver instance running locally.
	caps := selenium.Capabilities{"browserName": "chrome"}
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		panic(err)
	}
	defer wd.Quit()

	// Navigate to the url interface.
	if err := wd.Get("https://zh.surveymonkey.com/r/EmployeeHealthcheck"); err != nil {
		panic(err)
	}
	//click
	list := strings.Split(liststr, ",")
	click(wd, list)

	// Click the emplid button.
	btn, err := wd.FindElement(selenium.ByCSSSelector, "#\\36 83674383")
	if err != nil {
		panic(err)
	}
	if err := btn.SendKeys(emplid); err != nil {
		panic(err)
	}

	// Click the temperature button.
	//get random temperature
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	temperature := math.Round((r1.Float64()*(tempMax-tempMin)+tempMin)*1000) / 1000
	tempStr := fmt.Sprintf("%.2f", temperature)
	btn, err = wd.FindElement(selenium.ByCSSSelector, "#\\36 83674384")
	if err != nil {
		panic(err)
	}
	if err = btn.SendKeys(tempStr); err != nil {
		panic(err)
	}

	// Click the submit button.
	btn, err = wd.FindElement(selenium.ByCSSSelector, `button[type="submit"]`)
	if err != nil {
		panic(err)
	}
	if err = btn.Click(); err != nil {
		panic(err)
	}

	time.Sleep(100 * time.Second)
}
