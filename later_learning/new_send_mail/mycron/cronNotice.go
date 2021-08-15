package mycron

import (
	"github.com/robfig/cron/v3"
	"github.com/wonderivan/logger"
	"time"
)

type Fn func(string)

func CronNotice(info, timeStr, parameter string, fn Fn) {
	c := cron.New(cron.WithSeconds())
	//定时任务
	spec := timeStr //每天定点执行一次
	c.AddFunc(spec, func() {
		logger.Info(info + " : %v\n", time.Now())
		//以下为定时执行的操作
		fn(parameter)
	})
	c.Start()
}
