package data_svr

import (
	"fmt"
	"log"
	"os"
	"data_service/httpdb"
	"data_service/sendmail"
	"data_service/svr_common"
)

var Hdb httpdb.HttpDb
var SMail sendmail.Mailer

func InitConf() {
	// 通过配置文件配置

	logFile, err := os.OpenFile("/data/qb/data_service/logs/dataservice.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("os.OpenFile error :", err)
		return
	}
	// 设置输出位置 ，里面有锁进行控制
	log.SetOutput(logFile)

	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
}

func InitSvr() {
	// 初始化 httpdb 服务，抓取获取
	Hdb = httpdb.HttpDb{}

	// 初始化邮件服务
	SMail = sendmail.Mailer{
		Auth: sendmail.Auth{
			MailAddrAndPort: svr_common.MailAddrAndPort,
			Account:         svr_common.Account,
			Token:           svr_common.Token,
			MailAddr:        svr_common.MailAddr,
			Port:            587,
		},
		Sender: svr_common.Sender,
		To:     svr_common.To,
		Cc:     svr_common.Cc,
	}
}
