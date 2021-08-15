package main

import (
	"fmt"
	"log"
	"time"
	"xmt/xmtmq"
)

/*
RabbimtMQ routing 模式 案例
应用场景：从系统的代码逻辑中获取对应的功能字符串,将消息任务扔到对应的队列中业务场景,例如处理错误，处理特定消息等
生产消息
*/

func main() {

	log.SetFlags(log.Llongfile | log.Ltime | log.Ldate)

	rbt1 := &xmtmq.RabbitMQ{
		Exchange: "xmtPubEx2",
		Key: "xmt1",
		QueueName: "Routingqueuexmt1",
		MQUrl:     "amqp://guest:guest@127.0.0.1:5672/xmtmq",
	}

	xmtmq.NewRabbitMQ(rbt1)
	rbt1.Init()


	rbt2 := &xmtmq.RabbitMQ{
		Exchange: "xmtPubEx2",
		Key: "xmt2",
		QueueName: "Routingqueuexmt2",
		MQUrl:     "amqp://guest:guest@127.0.0.1:5672/xmtmq",
	}

	xmtmq.NewRabbitMQ(rbt2)
	rbt2.Init()


	var index = 0

	for {
		rbt1.ProduceRouting([]byte(fmt.Sprintf("hello wolrd xmt1  %d ", index)))
		log.Println("发送成功xmt1  ", index)

		rbt2.ProduceRouting([]byte(fmt.Sprintf("hello wolrd xmt2  %d ", index)))
		log.Println("发送成功xmt2  ", index)


		index++
		time.Sleep(1 * time.Second)
	}


	xmtmq.RabbitMQFree(rbt1)
	xmtmq.RabbitMQFree(rbt2)

}
