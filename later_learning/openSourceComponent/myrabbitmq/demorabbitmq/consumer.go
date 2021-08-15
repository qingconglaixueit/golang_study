package main

import "singlerabbitmq/xmtrbtmq"

func main() {
	rabbitmq := xmtrbtmq.NewRabbitMQSimple("xmtqueue")
	rabbitmq.ConsumeSimple()
}
