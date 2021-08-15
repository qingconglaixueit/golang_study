package main

import (
	"context"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/clientv3/concurrency"
	"log"
)

func main (){

	// 设置 log 参数 ，打印当前时间 和 当前行数
	log.SetFlags(log.Ltime | log.Llongfile)

	// ETCD 默认端口号是 2379
	// 使用 ETCD 的 clientv3 包
	// Endpoints 需填入 url 列表
	client, err := clientv3.New(clientv3.Config{Endpoints: []string{"/name"}})
	if err != nil {
		log.Printf("connect to etcd error : %v\n", err)
		return
	}
	defer client.Close()

	// 创建第一个 会话
	session1, err := concurrency.NewSession(client)
	if err != nil {
		log.Printf("concurrency.NewSession 1 error : %v\n", err)
		return
	}
	defer session1.Close()
	// 设置锁
	myMu1 := concurrency.NewMutex(session1, "/lock")

	// 创建第二个 会话
	session2, err := concurrency.NewSession(client)
	if err != nil {
		log.Printf("concurrency.NewSession 2 error : %v\n", err)
		return
	}
	defer session2.Close()
	// 设置锁
	myMu2 := concurrency.NewMutex(session2, "/lock")

	// 会话s1获取锁
	if err := myMu1.Lock(context.TODO()); err != nil {
		log.Printf("myMu1.Lock error : %v\n", err)
		return
	}
	log.Println("Get session1 lock ")


	m2Locked := make(chan struct{})
	go func() {
		defer close(m2Locked)
		// 如果加锁不成功会阻塞，知道加锁成功为止
		// 这里是使用一个通道的方式来通信
		// 当 myMu2 能加锁成功，说明myMu1 解锁成功
		// 当 myMu2 加锁成功的时候，会关闭 通道
		// 关闭通道，从通道中读出来的就是nil
		if err := myMu2.Lock(context.TODO()); err != nil {
			log.Printf("myMu2.Lock error : %v\n", err)
			return
		}
	}()

	// 解锁
	if err := myMu1.Unlock(context.TODO()); err != nil {
		log.Printf("myMu1.Unlock error : %v\n", err)
		return
	}
	log.Println("Release session1 lock ")

	// 读取到nil
	<-m2Locked

	log.Println("Get session2 lock")
}
