package data_svr

import (
	"data_service/svr_common"
	"fmt"
	"log"
)

type Greeting struct{}

// 晚上的问候
func (d Greeting) SendInfo(nightType string) {

	var bStr string
	switch nightType {
	case svr_common.Night:
		bStr = fmt.Sprintf("<h3>%s</h3>","好晚了，早点休息了, 别忘了喝点水再休息")
	case svr_common.Night2:
		bStr = fmt.Sprintf("<h3>%s</h3>", "要上床了, 再不睡就要变熊猫了")
	}

	makeData(&SMail, bStr,svr_common.NightSub)

	if err := SMail.SendMail(); err != nil {
		log.Println("SMail.SendMail error: ", err)
	}
}
