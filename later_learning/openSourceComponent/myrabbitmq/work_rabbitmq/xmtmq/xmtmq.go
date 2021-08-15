package xmtmq

import (
	"github.com/streadway/amqp"
	"log"
)

// work 模式
// 定义 RabbitMQ 的数据结构
// go get github.com/streadway/amqp

type RabbitMQ struct {
	conn      *amqp.Connection // 连接
	channel   *amqp.Channel    // 通道
	QueueName string           // 队列名
	Exchange  string           // 交换机
	Key       string           // 路由键
	MQUrl     string           // MQ的虚拟机地址
}

// New 一个 RabbitMQ
func NewRabbitMQ(rbt *RabbitMQ) {
	if rbt == nil || rbt.QueueName == "" || rbt.MQUrl == "" {
		log.Panic("please check QueueName,Exchange,MQUrl ...")
	}

	conn, err := amqp.Dial(rbt.MQUrl)
	if err != nil {
		log.Panicf("amqp.Dial error : %v", err)
	}
	rbt.conn = conn

	channel, err := rbt.conn.Channel()
	if err != nil {
		log.Panicf("rbt.conn.Channel error : %v", err)
	}
	rbt.channel = channel
}

func RabbitMQFree(rbt *RabbitMQ) {
	if rbt == nil {
		log.Printf("rbt is nil,free failed")
		return
	}

	rbt.channel.Close()
	rbt.conn.Close()

}

func (rbt *RabbitMQ) Init() {
	// 申请队列
	_, err := rbt.channel.QueueDeclare(
		rbt.QueueName, // 队列名
		true,          // 是否持久化
		false,         // 是否自动删除
		false,         // 是否排他
		false,         // 是否阻塞
		nil,           // 其他参数
	)
	if err != nil {
		log.Printf("rbt.channel.QueueDeclare error : %v", err)
		return
	}
}

// 生产消息
func (rbt *RabbitMQ) Produce(data []byte) {
	// 向队列中加入数据
	err := rbt.channel.Publish(
		rbt.Exchange,  // 交换机
		rbt.QueueName, // 队列名
		false,         // 若为true，根据自身exchange类型和routekey规则无法找到符合条件的队列会把消息返还给发送者
		false,         // 若为true，当exchange发送消息到队列后发现队列上没有消费者，则会把消息返还给发送者
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        data,
		},
	)
	if err != nil {
		log.Printf("rbt.channel.Publish error : %v", err)
		return
	}
	return
}

// 消费消息
func (rbt *RabbitMQ) Consume() {

	// 1、消费数据
	msg, err := rbt.channel.Consume(
		rbt.QueueName, // 队列名
		"xmt",         // 消费者的名字
		true,          // 是否自动应答
		false,         // 是否排他
		false,         // 若为true，表示 不能将同一个Conenction中生产者发送的消息传递给这个Connection中 的消费者
		false,         // 是否阻塞
		nil,           // 其他属性
	)

	if err != nil {
		log.Printf("rbt.channel.Consume error : %v", err)
		return
	}

	for data := range msg {
		log.Printf("received data is %v", string(data.Body))
	}
}
