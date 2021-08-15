package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"log"
	"time"
)

//func main() {
//
//	// 设置 log 参数 ，打印当前时间 和 当前行数
//	log.SetFlags(log.Ltime | log.Llongfile)
//
//	 // ETCD 默认端口号是 2379
//	// 使用 ETCD 的 clientv3 包
//	client, err := clientv3.New(clientv3.Config{
//		Endpoints:   []string{"127.0.0.1:2379"},
//		//超时时间 10 秒
//		DialTimeout: 10 * time.Second,
//	})
//
//	if err != nil {
//		log.Printf("connect to etcd error : %v\n", err)
//		return
//	}
//
//	log.Printf("connect to etcd successfully ...")
//	// defer 最后关闭 连接
//	defer client.Close()
//
//	// PUT KEY 为 name , value 为 xiaomotong
//	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
//	_, err = client.Put(ctx, "name", "xiaomotong")
//	cancel()
//	if err != nil {
//		log.Printf("PUT key to etcd error : %v\n", err)
//		return
//	}
//
//	// 获取ETCD 的KEY
//	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
//	resp, err := client.Get(ctx, "name")
//	cancel()
//	if err != nil {
//		log.Printf("GET key-value from etcd error : %v\n", err)
//		return
//	}
//
//	for _, ev := range resp.Kvs {
//		log.Printf("%s : %s\n", ev.Key, ev.Value)
//	}
//}

func main() {

	// 设置 log 参数 ，打印当前时间 和 当前行数
	log.SetFlags(log.Ltime | log.Llongfile)
	// ETCD 默认端口号是 2379
	// 使用 ETCD 的 clientv3 包
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 10 * time.Second,
	})
	if err != nil {
		log.Printf("connect to etcd error : %v\n", err)
		return
	}

	log.Printf("connect to etcd successfully ...")

	defer client.Close()

	// 我们创建一个 20秒钟的租约
	resp, err := client.Grant(context.TODO(), 20)
	if err != nil {
		log.Printf("client.Grant error : %v\n", err)
		return
	}

	// 20秒钟之后, /name 这个key就会被移除
	_, err = client.Put(context.TODO(), "/name", "xiaomotong", clientv3.WithLease(resp.ID))
	if err != nil {
		log.Printf("client.Put error : %v\n", err)
		return
	}


	// 这个key  name ，将永久被保存
	ch, kaerr := client.KeepAlive(context.TODO(), resp.ID)
	if kaerr != nil {
		log.Fatal(kaerr)
	}
	for {
		ka := <-ch
		fmt.Println("ttl:", ka.TTL)
	}
}