package xmtmq

import (
	"github.com/streadway/amqp"
	"log"
	"math/rand"
)

// rpc 模式
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

// 生产消息

func (rbt *RabbitMQ) Produce(data []byte) {

	// 申请队列
	q, err := rbt.channel.QueueDeclare(
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

	err = rbt.channel.Qos(1, 0, false)
	if err != nil {
		log.Printf("rbt.channel.Qos error : %v", err)
		return
	}

	d, err := rbt.channel.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil)
	if err != nil {
		log.Printf("rbt.channel.Consume error : %v", err)
		return
	}

	for msg := range d {
		log.Println("received msg is  ", string(msg.Body))
		err := rbt.channel.Publish(
			"",
			msg.ReplyTo,
			false,
			false,
			amqp.Publishing{
				ContentType:   "test/plain",
				CorrelationId: msg.CorrelationId,
				Body:          data,
			})
		if err != nil {
			log.Printf("rbt.channel.Publish error : %v", err)
			return
		}
		msg.Ack(false)
		log.Println("svr response ok ")
	}

	return
}
func randomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(rand.Intn(l))
	}
	return string(bytes)
}

// 消费消息
func (rbt *RabbitMQ) Consume() {

	// 申请队列
	q, err := rbt.channel.QueueDeclare(
		"",    // 队列名
		true,  // 是否持久化
		false, // 是否自动删除
		false, // 是否排他
		false, // 是否阻塞
		nil,   // 其他参数
	)
	if err != nil {
		log.Printf("rbt.channel.QueueDeclare error : %v", err)
		return
	}

	// 消费数据
	msg, err := rbt.channel.Consume(
		q.Name, // 队列名
		"xmt",  // 消费者的名字
		true,   // 是否自动应答
		false,  // 是否排他
		false,  // 若为true，表示 不能将同一个Conenction中生产者发送的消息传递给这个Connection中 的消费者
		false,  // 是否阻塞
		nil,    // 其他属性
	)
	if err != nil {
		log.Printf("rbt.channel.Consume error : %v", err)
		return
	}
	id := randomString(32)
	err = rbt.channel.Publish(
		"",
		rbt.QueueName,
		false,
		false,
		amqp.Publishing{
			ContentType:   "test/plain",
			CorrelationId: id,
			ReplyTo:       q.Name,
			Body:          []byte("321"),
		})
	if err != nil {
		log.Printf("rbt.channel.Publish error : %v", err)
		return
	}

	for data := range msg {
		log.Printf("received data is %v", string(data.Body))
	}
}
