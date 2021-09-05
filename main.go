package main

import (
	"fmt"
	"log"
	"os"
	"time"

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
	list := []string{"#\\36 83674386_4495696088", "#\\36 83674383", "#\\36 83674388_4495696090", "#\\36 83674400_4495696174", "#\\36 83674393_4495696115",
		"#\\36 83711504_4495952679", "#\\36 83674394_4495717677", "#\\36 83674398_4495718982", "#\\36 83674395_4495696119", "#\\36 83674397_4495696166", "#\\36 83674385_4495696080"}
	click(wd, list)

	// Click the emplid button.
	btn, err := wd.FindElement(selenium.ByCSSSelector, "#\\36 83674383")
	if err != nil {
		panic(err)
	}
	if err := btn.SendKeys("092844"); err != nil {
		panic(err)
	}

	// Click the temperature button.
	btn, err = wd.FindElement(selenium.ByCSSSelector, "#\\36 83674384")
	if err != nil {
		panic(err)
	}
	if err = btn.SendKeys("36"); err != nil {
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
