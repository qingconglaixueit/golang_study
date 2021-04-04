package main

import (
	"github.com/wonderivan/logger"
	"myMail/httpdb"
	"time"
)

func TimeWeatherM() {
	logger.Info("start TimeWeatherM...")
	for {
		now := time.Now()                                                                     //获取当前时间，放到now里面，要给next用  
		next := now.Add(time.Hour * 24)                                                       //通过now偏移24小时
		next = time.Date(next.Year(), next.Month(), next.Day(), 7, 25, 0, 0, next.Location()) //获取下一个凌晨的日期
		t := time.NewTimer(next.Sub(now))                                                     //计算当前时间到next的时间间隔，设置一个定时器
		<-t.C
		logger.Info("早晨天气预报: %v\n", time.Now())
		//以下为定时执行的操作
		httpdb.SendWeather(httpdb.DailuWeather)
	}
}

func TimeWeatherN() {
	logger.Info("start TimeWeatherN...")
	for {
		now := time.Now()                                                                      //获取当前时间，放到now里面，要给next用  
		next := now.Add(time.Hour * 24)                                                        //通过now偏移24小时
		next = time.Date(next.Year(), next.Month(), next.Day(), 23, 20, 0, 0, next.Location()) //获取下一个凌晨的日期
		t := time.NewTimer(next.Sub(now))                                                      //计算当前时间到next的时间间隔，设置一个定时器
		<-t.C
		logger.Info("现在天气预报: %v\n", time.Now())
		//以下为定时执行的操作
		httpdb.SendWeather(httpdb.DailuWeather)
	}
}

func TimeMorning() {
	logger.Info("start TimeMorning...")
	for {
		now := time.Now()                                                                     //获取当前时间，放到now里面，要给next用  
		next := now.Add(time.Hour * 24)                                                       //通过now偏移24小时
		next = time.Date(next.Year(), next.Month(), next.Day(), 9, 20, 0, 0, next.Location()) //获取下一个凌晨的日期
		t := time.NewTimer(next.Sub(now))                                                     //计算当前时间到next的时间间隔，设置一个定时器
		<-t.C
		logger.Info("早上好: %v\n", time.Now())
		//以下为定时执行的操作
		httpdb.SendDaily(httpdb.Morning)
	}
}

func TimeAfternoon() {
	logger.Info("start TimeAfternoon...")
	for {
		now := time.Now()                                                                      //获取当前时间，放到now里面，要给next用  
		next := now.Add(time.Hour * 24)                                                        //通过now偏移24小时
		next = time.Date(next.Year(), next.Month(), next.Day(), 12, 55, 0, 0, next.Location()) //获取下一个凌晨的日期
		t := time.NewTimer(next.Sub(now))                                                      //计算当前时间到next的时间间隔，设置一个定时器
		<-t.C
		logger.Info("中午了: %v\n", time.Now())
		//以下为定时执行的操作
		httpdb.SendDaily(httpdb.Afternoon)
	}
}

func TimeNight() {
	logger.Info("start TimeNight...")
	for {
		now := time.Now()                                                                      //获取当前时间，放到now里面，要给next用  
		next := now.Add(time.Hour * 24)                                                        //通过now偏移24小时
		next = time.Date(next.Year(), next.Month(), next.Day(), 20, 20, 0, 0, next.Location()) //获取下一个凌晨的日期
		t := time.NewTimer(next.Sub(now))                                                      //计算当前时间到next的时间间隔，设置一个定时器
		<-t.C
		logger.Info("晚上了: %v\n", time.Now())
		//以下为定时执行的操作
		httpdb.SendDaily(httpdb.Night)
	}
}

//按摩小秘密
func TimeMassage() {
	logger.Info("start 按摩小秘密...")
	for {
		now := time.Now()                                                                      //获取当前时间，放到now里面，要给next用  
		next := now.Add(time.Hour * 24)                                                        //通过now偏移24小时
		next = time.Date(next.Year(), next.Month(), next.Day(), 22, 30, 0, 0, next.Location()) //获取下一个凌晨的日期
		t := time.NewTimer(next.Sub(now))                                                      //计算当前时间到next的时间间隔，设置一个定时器
		<-t.C
		logger.Info("按摩小秘密: %v\n", time.Now())
		//以下为定时执行的操作
		httpdb.SendDaily(httpdb.Night2)
	}
}

func main() {

	// 通过配置文件配置
	logger.SetLogger("./log.json")
	logger.Info(" ------------ start  main ------------")

	go TimeWeatherM()
	go TimeWeatherN()
	go TimeMorning()
	go TimeAfternoon()
	go TimeNight()
	go TimeMassage()

	select {}

	//resp, err := http.Get("https://www.cnblogs.com/xiaomotong/")
	//if err != nil {
	//	fmt.Printf("Post err %v", err)
	//	return
	//}
	//defer resp.Body.Close()
	//
	//b, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	fmt.Printf("ReadAll err %v", err)
	//	return
	//}
	////fmt.Println(string(b))
	//
	//htmlContent := string(b)
	//
	//city ,err := httpdb.GetSpecialData(htmlContent, ".post-view-count")
	//log.Printf(city)
	//
	//city ,err = httpdb.GetSpecialData(htmlContent, "a span")
	//log.Printf(city)
}
