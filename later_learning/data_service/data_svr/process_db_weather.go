package data_svr

import (
	"data_service/svr_common"
	"fmt"
	"log"
	"strings"
	"time"
)

type Weather struct{}

// 发送每日一句和天气预报
func (d Weather) SendInfo(link string) {
	weather := GetWeatherInfo(link)
	dailyInfo := GetDailyInfo()

	makeData(&SMail, fmt.Sprintf("%s<br/>%s",
		weather, dailyInfo),svr_common.DailySub)

	if err := SMail.SendMail(); err != nil {
		log.Println("SMail.SendMail error: ", err)
	}
	return
}

func GetDailyInfo() string {
	//获取当前年月日
	res := strings.Split(time.Now().String(), " ")
	log.Printf("C(time.Now().String() ==  %s\n", res[0])

	//queryUrl := fmt.Sprintf(link, res[0])
	log.Printf("Chrome visit page %s\n", svr_common.DailyLink)
	htmlContent, err := Hdb.GetHttpHtmlContent(svr_common.DailyLink, "body > div.screen > div.banner > div.swiper-container-place > div > div.swiper-slide.swiper-slide-0.swiper-slide-visible.swiper-slide-active > a.item.item-big > div.item-bottom", `document.querySelector("body")`)
	if err != nil {
		log.Printf("GetHttpHtmlContent err : %v", err)
		return ""
	}

	//log.Printf(htmlContent)

	strZh, err := Hdb.GetSpecialData(htmlContent, ".chinese")
	if err != nil {
		log.Printf("GetSpecialData err : %v", err)
		return ""
	}

	log.Println(strZh)

	strEn, err := Hdb.GetSpecialData(htmlContent, ".english")
	if err != nil {
		log.Printf("GetSpecialData err : %v", err)
		return ""
	}
	log.Println(strEn)

	return fmt.Sprintf("<h3>%s</h3><h4>%s</h4>", strZh, strEn)
}

// 天气预报信息
func GetWeatherInfo(link string) string {
	htmlContent, err := Hdb.GetHttpHtmlContent(link, "body > div.weatherbox > div > div.left > dl > dd.weather > span", `document.querySelector(".weather_info")`)
	if err != nil {
		log.Printf("GetHttpHtmlContent err : %v", err)
		return ""
	}

	//log.Printf(htmlContent)

	//src, err := Hdb.GetSpecialAttrData(htmlContent, "dt img")
	//log.Println(src)
	//filename := GetPic(src)
	// 城市天气
	city, err := Hdb.GetSpecialData(htmlContent, "h1")
	log.Println(city)
	// 当前气温
	weatherNum, err := Hdb.GetSpecialData(htmlContent, ".weather .now")
	log.Println(weatherNum)
	// 气温范围
	weather, err := Hdb.GetSpecialData(htmlContent, ".weather span")
	log.Println(weather)

	return fmt.Sprintf("<h3>%s</h3><h3>当前气温：%s<h3/><h3>今日气温范围： %s <h3/>", city, weatherNum, weather)
}
