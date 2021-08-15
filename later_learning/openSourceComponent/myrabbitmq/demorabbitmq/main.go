package main

/*

my_rabbitmq
rabbitmq 的系统学习

*/

import (
	"fmt"
	"log"
	"singlerabbitmq/xmtrbtmq"
	"time"
)

func main() {

	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)

	rabitmq := xmtrbtmq.NewRabbitMQSimple("xmtqueue")

	var index = 1
	for {
		rabitmq.PublishSimple(fmt.Sprintf("xmt hello world -- %d", index))
		fmt.Println("发送成功!   ", index)

		index++

		time.Sleep(2 * time.Second)

	}

}
