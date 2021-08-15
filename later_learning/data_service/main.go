package main

import (
	"data_service/data_svr"
	"data_service/mycron"
	"log"
	"data_service/svr_common"
)

func main() {
	// 初始化服务
	data_svr.InitSvr()
	data_svr.InitConf()

	// 执行任务
	log.Println(" ------------ start  main ------------")

	// ------------------------------ 每天天气预报 和 每日一句 -------------------------------
	w := data_svr.Weather{}
	m := mycron.MyCron{
		TimeStr:   "00 00 9,21 * * ?",
		Parameter: svr_common.GuangzhouLink,
		Fn:        w.SendInfo,
	}
	m.CronNotice()
	// ------------------------------ 每天天气预报 和 每日一句 -------------------------------

	// ------------------------------ 发布雪球热榜 -------------------------------
	f := data_svr.Finance{}
	// 雪球 沪深热榜
	m5 := mycron.MyCron{
		TimeStr:   "00 10 9,12,16,19,21 * * ?",
		Parameter: svr_common.XueQiu,
		Fn:        f.SendInfo,
	}
	m5.CronNotice()
	// ------------------------------ 发布雪球热榜 -------------------------------

	// ------------------------------ 每日晚间问候 -------------------------------
	g := data_svr.Greeting{}
	m2 := mycron.MyCron{
		TimeStr:   "00 30 21 * * ?",
		Parameter: svr_common.Night,
		Fn:        g.SendInfo,
	}
	m2.CronNotice()
	// ------------------------------ 每日晚间问候 -------------------------------

	// ------------------------------ 每日晚间问候2 -------------------------------
	m3 := mycron.MyCron{
		TimeStr:   "00 20 22 * * ?",
		Parameter: svr_common.Night2,
		Fn:        g.SendInfo,
	}
	m3.CronNotice()
	// ------------------------------ 每日晚间问候2 -------------------------------




	log.Println(" ------------ wait next cron ------------")

	select {}
}
