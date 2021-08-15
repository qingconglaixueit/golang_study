package main

import (
	"context"
	"log"

	"mygrpc.com/pb"

	"time"

	"google.golang.org/grpc"
)

func main() {
	// 连接grpc服务
	conn, err := grpc.Dial(":8888", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	// 很关键
	defer conn.Close()

	// 初始化客户端
	c := pb.NewLoveClient(conn)

	// 发起请求
	var i int
	for i = 0; i < 10; i++ {
		response, err := c.Confession(context.Background(), &pb.Request{Name: "linnux kernel"})
		if err != nil {
			log.Fatal(err)
		}
		log.Println(response.Result)
		time.Sleep(1000 * time.Millisecond)
	}

}
