package business

//toda regra de negócio
//getTemperature

import (
	"log"

	"time"

	"go-gin-api/model"
	//"go-gin-api/internal"
	"github.com/sclevine/agouti"
)

func GetTemperatures() model.Temperature {

	morningTemp := getItem(`//body/div/main/div[2]/main/div[3]/section/div/ul/li[1]/a/div/span[@data-testid="TemperatureValue"]`)

	log.Printf("The temperature in the morning is %v, as of %v ", morningTemp, time.Now())

	afternoonTemp := getItem(`//body/div/main/div[2]/main/div[3]/section/div/ul/li[2]/a/div/span[@data-testid="TemperatureValue"]`)
	log.Printf("The temperature in the afternoon is %v, as of %v ", afternoonTemp, time.Now())

	eveningTemp := getItem(`//body/div/main/div[2]/main/div[3]/section/div/ul/li[3]/a/div/span[@data-testid="TemperatureValue"]`)
	log.Printf("The temperature in the evening is %v, as of %v ", eveningTemp, time.Now())

	nightTemp := getItem(`//body/div/main/div[2]/main/div[3]/section/div/ul/li[4]/a/div/span[@data-testid="TemperatureValue"]`)
	log.Printf("The temperature at night is %v, as of %v", nightTemp, time.Now())

	//var temp model.Temperature = model.Temperature{morningTemp, afternoonTemp, eveningTemp, nightTemp }
	temp := model.Temperature{morningTemp, afternoonTemp, eveningTemp, nightTemp}
	//fmt.Println(temp.MorningT)

	return temp

}

func getItem(s string) string {
	page, driver := setUpAgouti()

	if err := page.Navigate("https://weather.com/weather/today/l/93f6a6719a5327824688b56f52a87e64438d9947fd2c3da96b580ba43de24a33"); err != nil {
		log.Fatal("Failed to navigate:", err)
	}

	value, err := page.FindByXPath(s).Text()

	timeForWait := time.Second * 1 //5

	time.Sleep(timeForWait)

	if err != nil {
		log.Fatal("Failed to find item", err)
	}

	//log.Println(value)

	if err := driver.Stop(); err != nil {
		log.Fatal("Failed to close page and stop webdriver:", err)
	}

	return value

}

func setUpAgouti() (*agouti.Page, *agouti.WebDriver) {
	driver := agouti.ChromeDriver()

	if err := driver.Start(); err != nil {
		log.Fatal("Failed to start driver:", err)
	}

	page, err := driver.NewPage()
	if err != nil {
		log.Fatal("Failed to open page:", err)
	}

	return page, driver

}
