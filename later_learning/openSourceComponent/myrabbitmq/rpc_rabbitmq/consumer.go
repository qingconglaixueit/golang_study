package main

import (
	"log"
	"math/rand"
	"time"
	"xmt/xmtmq"
)

func main() {

	log.SetFlags(log.Llongfile | log.Ltime | log.Ldate)
	rand.Seed(time.Now().UTC().UnixNano())
	rbt := &xmtmq.RabbitMQ{
		QueueName: "xmtqueue",
		MQUrl:     "amqp://guest:guest@127.0.0.1:5672/xmtmq",
	}

	xmtmq.NewRabbitMQ(rbt)
	rbt.Consume()
}
