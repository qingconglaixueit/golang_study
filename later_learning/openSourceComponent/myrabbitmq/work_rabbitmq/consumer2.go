package main

import (
	"log"
	"xmt/xmtmq"
)

func main() {

	log.SetFlags(log.Llongfile | log.Ltime | log.Ldate)

	rbt := &xmtmq.RabbitMQ{
		QueueName: "xmtqueue",
		MQUrl:     "amqp://guest:guest@127.0.0.1:5672/xmtmq",
	}

	xmtmq.NewRabbitMQ(rbt)
	rbt.Consume()
	xmtmq.RabbitMQFree(rbt)
}
