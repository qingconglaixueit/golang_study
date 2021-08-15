package main

import (
	"log"
	"xmt/xmtmq"
)

func main() {

	log.SetFlags(log.Llongfile | log.Ltime | log.Ldate)

	rbt := &xmtmq.RabbitMQ{
		Exchange: "xmtPubEx3",
		Key: "#.xmt1",
		QueueName: "Topicxmt1",
		MQUrl:     "amqp://guest:guest@127.0.0.1:5672/xmtmq",
	}

	xmtmq.NewRabbitMQ(rbt)
	rbt.ConsumeTopicMsg()
	xmtmq.RabbitMQFree(rbt)
}
