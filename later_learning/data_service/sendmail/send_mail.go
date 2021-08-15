package sendmail

import (
	"gopkg.in/gomail.v2"
	"log"
)

type Auth struct {
	MailAddrAndPort string
	Account         string
	Token           string
	MailAddr        string
	Port            int
}

// 邮件服务
type Mailer struct {
	Auth
	Call    string // 称呼
	Sub     string // 主题
	BodyStr string // 内容
	Sender  string
	To      string
	Cc      string
	Attach  string // 附件
}

func (data Mailer) SendMail() error {
	m := gomail.NewMessage()
	//发送人
	m.SetHeader("From", data.Sender)
	//接收人
	m.SetHeader("To", data.To)
	//抄送人
	m.SetAddressHeader("Cc", data.Cc, data.Call)
	//主题
	m.SetHeader("Subject", data.Sub)
	//内容
	m.SetBody("text/html", data.BodyStr)
	//附件
	//m.Attach("./myIpPic.png")

	//拿到token，并进行连接,第4个参数是填授权码
	d := gomail.NewDialer(data.MailAddr, data.Port, data.Sender, data.Token)

	// 发送邮件
	if err := d.DialAndSend(m); err != nil {
		log.Printf("DialAndSend err %v:", err)
		return err
	}
	log.Printf("send mail success\n")
	return nil
}

//func (m Mailer) SendMail() error {
//	em := email.NewEmail()
//	// 设置 sender 发送方 的邮箱 ， 此处可以填写自己的邮箱
//	em.From = m.Sender
//
//	// 设置 receiver 接收方 的邮箱  此处也可以填写自己的邮箱， 就是自己发邮件给自己
//	em.To = []string{m.To}
//
//	// 抄送
//	em.Cc = []string{m.Cc}
//
//	// 设置主题
//	em.Subject = m.Sub
//
//	// 简单设置文件发送的内容，暂时设置成纯文本
//	em.HTML = []byte(m.BodyStr)
//
//	//设置服务器相关的配置
//	err := em.Send(m.MailAddrAndPort, smtp.PlainAuth("", m.Account, m.Token, m.MailAddr))
//	if err != nil {
//		return err
//	}
//	log.Println("send successfully ... ")
//	return nil
//}
