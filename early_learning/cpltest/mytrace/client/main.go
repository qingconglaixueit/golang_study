package main

import (
	"log"
	pb "myclient/protoc/hi" // 引入proto包

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

const (
	// Address gRPC服务地址
	Address = "127.0.0.1:9999"
)


func main() {
	log.SetFlags(log.Ltime | log.Llongfile)
	// TLS连接  记得把xxx改成你写的服务器地址

	var err error

	conn, err := grpc.Dial(Address,grpc.WithInsecure())
	if err != nil {
		grpclog.Fatalln(err)
	}

	defer conn.Close()

	// 初始化客户端
	c := pb.NewHiClient(conn)

	// 调用方法
	req := &pb.HiRequest{Name: "gRPC"}
	res, err := c.SayHi(context.Background(), req)
	if err != nil {
		log.Panicln(err)
	}

	log.Println(res.Message)
}
