package main

import (
	"fmt"
	"log"
	"myMail/httpdb"
	"os"
)

func initConf() {
	// 通过配置文件配置

	logFile, err := os.OpenFile("./mailLog.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("os.OpenFile error :", err)
		return
	}
	// 设置输出位置 ，里面有锁进行控制
	log.SetOutput(logFile)

	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
}

func main() {
	initConf()

	log.Println(" ------------ start  main ------------")

	mycron.CronNotice("早晨天气预报", "00 20 03 * * ?", httpdb.DailuWeather, httpdb.SendWeatherGz)
	mycron.CronNotice("早晨天气预报", "00 22 03 * * ?", httpdb.DailuWeather, httpdb.SendWeather)

	mycron.CronNotice("早上好，朋友圈的小可爱，注意安全哦  ", "00 35 03 * * ?", httpdb.Morning, httpdb.SendDaily)

	mycron.CronNotice("早晨天气预报", "00 10 06 * * ?", httpdb.DailuWeather, httpdb.SendWeatherGz)
	mycron.CronNotice("早晨天气预报", "00 12 06 * * ?", httpdb.DailuWeather, httpdb.SendWeather)


	mycron.CronNotice("晚上天气预报", "00 30 21 * * ?", httpdb.DailuWeather, httpdb.SendWeather)


	mycron.CronNotice("早上好", "00 30 08 * * ?", httpdb.Morning, httpdb.SendDaily)
	//mycron.CronNotice("中午了哟", "00 00 12 * * ?", httpdb.Afternoon, httpdb.SendDaily)
	mycron.CronNotice("好晚了哟", "00 00 22 * * ?", httpdb.Night, httpdb.SendDaily)
	mycron.CronNotice("再不睡就成熊猫了", "00 30 22 * * ?", httpdb.Night2, httpdb.SendDaily)

	log.Println(" ------------ wait next cron ------------")

	select {}
}
