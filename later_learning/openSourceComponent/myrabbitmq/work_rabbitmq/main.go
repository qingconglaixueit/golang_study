package main

import (
	"fmt"
	"log"
	"time"
	"xmt/xmtmq"
)

/*
RabbimtMQ work 模式 案例
应用场景:红包
生产消息
*/

func main() {

	log.SetFlags(log.Llongfile | log.Ltime | log.Ldate)

	rbt := &xmtmq.RabbitMQ{
		QueueName: "xmtqueue",
		MQUrl:     "amqp://guest:guest@127.0.0.1:5672/xmtmq",
	}

	xmtmq.NewRabbitMQ(rbt)
	rbt.Init()

	var index = 0

	for {
		rbt.Produce([]byte(fmt.Sprintf("hello wolrd %d ", index)))
		log.Println("发送成功 ", index)
		index++
		time.Sleep(1 * time.Second)
	}

	xmtmq.RabbitMQFree(rbt)

}
