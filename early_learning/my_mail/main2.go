package main

import (
	"fmt"
	"github.com/jordan-wright/email"
	"log"
	"net/smtp"
	"sync"
	"time"
)

func main() {

	// 简单设置l og 参数
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	// 创建有5 个缓冲的通道，数据类型是  *email.Email
	ch := make(chan *email.Email, 5)
	// 连接池
	p, err := email.NewPool(
		"smtp.qq.com:25",
		3,
		smtp.PlainAuth("", "502892037@qq.com", "muahhmqlnrmobjii", "smtp.qq.com"),
	)

	if err != nil {
		log.Fatal("email.NewPool error : ", err)
	}

	var wg sync.WaitGroup
	wg.Add(3)
	for i := 0; i < 3; i++ {
		go func() {
			defer wg.Done()
			for e := range ch {
				// 超时时间 10 秒
				err := p.Send(e, 10*time.Second)
				if err != nil {
					log.Printf( "p.Send error : %v , e = %v , i = %d\n", err , e, i)
				}
			}
		}()
	}

	for i := 0; i < 5; i++ {
		e := email.NewEmail()
		e.From = "xx <502892037@qq.com>"
		e.To = []string{"502892037@qq.com"}
		e.Subject = "test email.NewPool " + fmt.Sprintf("  the %d email",i)
		e.Text = []byte(fmt.Sprintf("test email.NewPool , the %d email !", i))
		ch <- e
	}

	// 关闭通道
	close(ch)
	// 等待子协程退出
	wg.Wait()
	log.Println("send successfully ... ")
}