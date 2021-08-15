package main

import (
	"log"
	"xmt/xmtmq"
)

func main() {

	log.SetFlags(log.Llongfile | log.Ltime | log.Ldate)

	rbt := &xmtmq.RabbitMQ{
		Exchange: "xmtPubEx3",
		Key:      "xmt2",
		QueueName: "Topicxmt2",
		MQUrl:    "amqp://guest:guest@127.0.0.1:5672/xmtmq",
	}

	xmtmq.NewRabbitMQ(rbt)
	rbt.ConsumeTopicMsg()
	xmtmq.RabbitMQFree(rbt)
}
