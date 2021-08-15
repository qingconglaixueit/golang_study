package data_svr

import (
	"data_service/sendmail"
	"data_service/svr_common"
	"fmt"
	"log"
	"time"
	"math/rand"
)

// 发送信息接口
type SendInfoEr interface {
	SendInfo(string)
}

// 处理节日标题
func processSpecialSub(data *sendmail.Mailer) bool {
	months := time.Now().Month()
	days := time.Now().Day()
	if months == time.March && (days == 7 || days == 8) {
		data.Sub = svr_common.GodlessSub
		return true
	}
	return false
}

// 处理节日称呼
func processSpecialCall(data *sendmail.Mailer) bool {
	months := time.Now().Month()
	days := time.Now().Day()
	// 3月7日，3月8日 女神节
	if months == time.March && (days == 7 || days == 8) {
		data.Call = svr_common.Godless
		return true
	}

	return false
}

// 处理节日内容
func processSpecialBody(data *sendmail.Mailer) bool {
	months := time.Now().Month()
	days := time.Now().Day()

	if months == time.March && (days == 7 || days == 8) {
		data.BodyStr = svr_common.GodlessBody
		return true
	}

	return false
}

// 处理标题
func makeSub(data *sendmail.Mailer, sub string) {

	if processSpecialSub(data) {
		return
	}

	data.Sub = data.Call + sub
}

// 处理内容
func makeBody(data *sendmail.Mailer, body string) {

	if processSpecialBody(data) {
		return
	}

	data.BodyStr = fmt.Sprintf("<h2>%s</h2>%s", data.Call, body)

}

// 处理称呼
func makeCall(data *sendmail.Mailer) {

	// 处理节日称呼
	if processSpecialCall(data) {
		return
	}

	// 处理通用称呼
	switch time.Now().Weekday() {
	case time.Sunday:
		data.Call = svr_common.CallName[0]
	case time.Monday:
		data.Call = svr_common.CallName[1]
	case time.Tuesday:
		data.Call = svr_common.CallName[2]
	case time.Wednesday:
		data.Call = svr_common.CallName[3]
	case time.Thursday:
		data.Call = svr_common.CallName[4]
	case time.Friday:
		data.Call = svr_common.CallName[5]
	case time.Saturday:
		data.Call = svr_common.CallName[6]
	default:
		data.Call = svr_common.CallName[rand.Int()%7]
	}

}

func makeData(data *sendmail.Mailer, body string, sub string) {
	if data == nil {
		log.Println("data is nil")
		return
	}
	// 标题
	makeSub(data, sub)
	// 称呼
	makeCall(data)
	// 内容
	makeBody(data, body)
	// 附件
	//if attach != "" {
	//	data.Attach = attach
	//}

}
