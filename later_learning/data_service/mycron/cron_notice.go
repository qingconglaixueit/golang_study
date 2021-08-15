package mycron

import (
	"github.com/robfig/cron/v3"
	"github.com/wonderivan/logger"
	"time"
)

type Fn func(string)

type MyCron struct {
	TimeStr   string
	Parameter string
	Fn        Fn
}

func (m MyCron) CronNotice() {
	c := cron.New(cron.WithSeconds())
	//定时任务
	spec := m.TimeStr //每天定点执行一次
	c.AddFunc(spec, func() {
		logger.Info(" starting cron :%v\n", time.Now())
		//以下为定时执行的操作
		m.Fn(m.Parameter)
	})
	c.Start()
}
