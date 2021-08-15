package svr_common

// daily call
const (
	Night  = "In the evening"
	Night2 = "go to bed"
)

// http url
const (
	//每日一句
	DailyLink = "http://news.iciba.com/"

	//广州天气预报
	GuangzhouLink = "https://www.tianqi.com/guangzhou/"
	// 达州天气预报
	DaZhouLink = "https://www.tianqi.com/dazhou/"

	// 雪球热股榜
	XueQiu = "https://xueqiu.com/"
)

// mail auth account
const (
	Account = "502892037"
	Sender  = "502892037@qq.com"  //发送人
	To      = "2569978958@qq.com" //接收人
	Cc      = "502892037@qq.com"  //抄送人

	//授权码
	Token           = "muahhmqlnrmobjii"
	MailAddr        = "smtp.qq.com"
	MailAddrAndPort = "smtp.qq.com:25"
)

// special call
const (
	Godless     = "ww小姐姐"
	GodlessBody = "女神节快乐哟，照顾好自己的情绪，爱自己，才能更好的爱更好的人哟！！"
	GodlessSub  = "女神节快乐"
)

// general call
var CallName = []string{
	"盆友圈的小可爱",
	"盆友圈的小可爱",
	"盆友圈的小可爱",
	"盆友圈的小可爱",
	"盆友圈的小可爱",
	"盆友圈的小可爱",
	"盆友圈的小可爱",
}

// mail sub
const (
	DailySub  = "今日天气 & 每日一句 唤醒真正野蛮生长的你 "
	NightSub  = " 晚间小问候 "
	XueQiuSub = "雪球热股榜"
)
