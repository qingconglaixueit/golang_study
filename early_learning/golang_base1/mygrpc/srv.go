package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"mygrpc.com/pb"
)

// 定义Love服务
type Love struct {
}

// 实现Love服务接口
func (l *Love) Confession(ctx context.Context, request *pb.Request) (*pb.Response, error) {
	resp := &pb.Response{}
	resp.Result = "the book name is " + request.Name
	return resp, nil
}

func main() {
	// 监听8888端口
	listen, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatal(err)
	}

	// 实例化grpc server
	s := grpc.NewServer()

	// 注册Love服务
	pb.RegisterLoveServer(s, new(Love))

	log.Println("Listen on 127.0.0.1:8888...")
	s.Serve(listen)
}
