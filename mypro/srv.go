package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"myagent/protoc/pop_agent"
	"net"
	"time"

	"google.golang.org/grpc"
)

type haha struct {
}

func (h *haha) BiStream(stream pop_agent.PopAgent_BiStreamServer) error {

	//go func() {
	connFlg := 0
	in, err := stream.Recv()
	if err == io.EOF {
		fmt.Println("read done")
		return nil
	}

	if err != nil {
		fmt.Println("error : ", err)
		return nil
	}

	if in.CmdName == "AgentConn" {
		fmt.Println("收到客户端的连接请求，正在回应客户端")
		fmt.Println("recv data from client : ", in)
		fmt.Printf("%T", in)
		fmt.Println("\n=============================\n")
		err = stream.Send(&pop_agent.StreamMsg{
			CmdName: "AgentConn",
			Code:    "1",
		})
		if err != nil {
			fmt.Println("send failed!!")
		}
		if in.Uuid == "222" {
			connFlg = 2
		}
	}
	if connFlg == 2 {
		for {
			fmt.Println("开始发送租户系统的增量信息")
			err = stream.Send(&pop_agent.StreamMsg{
				CmdName: "increment",
				Code:    "1",
			})
			if err != nil {
				fmt.Println("increment send failed!!")
			}
			time.Sleep(time.Second * time.Duration(1))
		}
	}

	return nil
}

func (h *haha) DUpMsgSync(ctx context.Context, req *pop_agent.StreamMsg) (*pop_agent.StreamMsg, error) {
	if req.CmdName == "UpControl" {
		fmt.Printf("\n收到上报数据\n")
		fmt.Printf("\n正在响应客户端...\n")
		resp := &pop_agent.StreamMsg{
			CmdName: "UpControl",
			Uuid:    "214214",
			Content: "hello client",
			Code:    "1",
		}
		fmt.Printf("\n响应完毕...\n")
		return resp, nil
	}

	if req.CmdName == "AllConfig" {
		fmt.Printf("\n收到上报请求全量数据\n")
		fmt.Printf("\n正在响应客户端...\n")
		resp := &pop_agent.StreamMsg{
			CmdName: "AllConfig",
			Uuid:    "214214",
			Content: "hello client",
			Code:    "1",
		}
		fmt.Printf("\n响应全量数据完毕...\n")
		return resp, nil
	}

	return nil, nil
}

func main() {
	listen, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()

	//注册服务
	pop_agent.RegisterPopAgentServer(server, &haha{})

	log.Println("start listen 8888...")
	err = server.Serve(listen)
	if err != nil {
		log.Fatalf("grpcServer.Serve err: %v", err)
	}
}
