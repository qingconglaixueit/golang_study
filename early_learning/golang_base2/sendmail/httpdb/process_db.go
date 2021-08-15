package httpdb

import (
	"fmt"
	"github.com/wonderivan/logger"
	"gopkg.in/gomail.v2"
	"log"
	"math/rand"
	"strings"
	"time"
)

type MyData struct {
	call    string // 名字程序
	sub     string // 主题
	bodyStr string // 内容
}

const (
	Morning      = "Good morning"
	Afternoon    = "lunch break"
	Night        = "In the evening"
	Night2       = "go to bed"
	DailuWeather = "Today's weather"
)

const (
	//每日一句
	link = "http://news.iciba.com/"
	//广州天气预报
	link2 = "https://www.tianqi.com/guangzhou/"

	sender = "502892037@qq.com"  //发送人
	toer   = "2569978958@qq.com" //接收人
	ccer   = "502892037@qq.com"  //抄送人

	//授权码
	token    = "muahhmqlnrmobjii"
	port     = 587
	mailAddr = "smtp.qq.com"
)
const (
	godless     = "ww小姐姐"
	godlessBody = "女神节快乐哟，照顾好自己的情绪，爱自己，才能更好的爱更好的人哟！！"
	godlessSub  = "女神节快乐"
)

var CallName = []string{"小佩奇", "小奋斗", "小可爱", "小加油", "小努力", "小变化", "小休息"}

func SendDaily(timeStr string) {
	if timeStr == Morning {
		//获取当前年月日
		res := strings.Split(time.Now().String(), " ")
		logger.Info("C(time.Now().String() ==  %s\n", res[0])

		//queryUrl := fmt.Sprintf(link, res[0])
		logger.Info("Chrome visit page %s\n", link)
		htmlContent, err := GetHttpHtmlContent(link, "body > div.screen > div.banner > div.swiper-container-place > div > div.swiper-slide.swiper-slide-0.swiper-slide-visible.swiper-slide-active > a.item.item-big > div.item-bottom", `document.querySelector("body")`)
		if err != nil {
			logger.Info("GetHttpHtmlContent err : %v", err)
			return
		}

		//log.Printf(htmlContent)

		strZh, err := GetSpecialData(htmlContent, ".chinese")
		if err != nil {
			logger.Info("GetSpecialData err : %v", err)
			return
		}

		logger.Info(strZh)

		strEn, err := GetSpecialData(htmlContent, ".english")
		if err != nil {
			log.Printf("GetSpecialData err : %v", err)
			return
		}
		logger.Info(strEn)

		bStr := fmt.Sprintf("<h3>%s,%s</h3><h4>%s</h4>", timeStr, strZh, strEn)

		data := &MyData{}
		MakeData(data, bStr)

		SendMail(data)
	} else if timeStr == Afternoon {
		bStr := fmt.Sprintf("<h3>%s,%s</h3><h4>%s</h4>", timeStr, "逛完路了吧,快回来休息了", "下午多喝点水，才能精神满满")

		data := &MyData{}
		MakeData(data, bStr)

		SendMail(data)
	} else if timeStr == Night {
		bStr := fmt.Sprintf("<h3>%s,%s</h3><h3>%s</h3>", timeStr, "早点处理好事情，早点回家", "注意安全哟，别忘了喝点水再回家")

		data := &MyData{}
		MakeData(data, bStr)

		SendMail(data)
	} else if timeStr == Night2 {
		bStr := fmt.Sprintf("<h3>%s , %s</h3><h3>%s</h3>", timeStr, "要上床了", "按摩自己的小姐姐哟")

		data := &MyData{}
		MakeData(data, bStr)

		SendMail(data)
	}
}

func SendWeather(nowWeather string) {

	htmlContent, err := GetHttpHtmlContent(link2, "body > div.weatherbox > div > div.left > dl > dd.weather > span", `document.querySelector(".weather_info")`)
	if err != nil {
		logger.Info("GetHttpHtmlContent err : %v", err)
		return
	}
	//log.Printf(htmlContent)

	city, err := GetSpecialData(htmlContent, ".name h2")
	logger.Info(city)

	weatherNum, err := GetSpecialData(htmlContent, ".weather .now")
	logger.Info(weatherNum)

	weather, err := GetSpecialData(htmlContent, ".weather span")
	logger.Info(weather)

	bStr := fmt.Sprintf("<h2>%s:</h2><h3>%s - %s - %s</h3>", nowWeather, city, weatherNum, weather)

	data := &MyData{}
	MakeData(data, bStr)

	SendMail(data)
}

func processSpecialDay(data *MyData) {
	if data == nil {
		logger.Warn("data is nil")
		return
	}
	months := time.Now().Month()
	days := time.Now().Day()

	if months == time.March && (days == 7 || days == 8) {
		data.call = godless
	}

}

func processSpecialBody(data *MyData, body string) {
	if data == nil {
		logger.Warn("data is nil")
		return
	}
	months := time.Now().Month()
	days := time.Now().Day()

	if months == time.March && (days == 7 || days == 8) {
		body = godlessBody
	}

	data.bodyStr = fmt.Sprintf("<h2>%s</h2>%s", data.call, body)

}

func processSpecialSub(data *MyData) {
	if data == nil {
		logger.Warn("data is nil")
		return
	}
	months := time.Now().Month()
	days := time.Now().Day()

	sub := "是我哦"

	if months == time.March && (days == 7 || days == 8) {
		sub = godlessSub
	}

	data.sub = data.call + sub
}

func MakeData(data *MyData, body string) {
	//处理名字
	switch time.Now().Weekday() {
	case time.Sunday:
		data.call = CallName[0]
	case time.Monday:
		data.call = CallName[1]
	case time.Tuesday:
		data.call = CallName[2]
	case time.Wednesday:
		data.call = CallName[3]
	case time.Thursday:
		data.call = CallName[4]
	case time.Friday:
		data.call = CallName[5]
	case time.Saturday:
		data.call = CallName[6]
	default:
		data.call = CallName[rand.Int()%7]
	}

	//特殊节日，特殊称呼
	processSpecialDay(data)
	//内容
	processSpecialBody(data, body)
	//标题
	processSpecialSub(data)

}

func SendMail(data *MyData) {
	if data == nil {
		logger.Info("data is nil")
		return
	}
	m := gomail.NewMessage()
	//发送人
	m.SetHeader("From", sender)
	//接收人
	m.SetHeader("To", toer)
	//抄送人
	m.SetAddressHeader("Cc", ccer, data.call)
	//主题
	m.SetHeader("Subject", data.sub)
	//内容
	m.SetBody("text/html", data.bodyStr)
	//附件
	//m.Attach("./myIpPic.png")

	//拿到token，并进行连接,第4个参数是填授权码
	d := gomail.NewDialer(mailAddr, port, sender, token)

	// 发送邮件
	if err := d.DialAndSend(m); err != nil {
		logger.Info("DialAndSend err %v:", err)
		return
	}
	logger.Info("send mail success\n")
}
