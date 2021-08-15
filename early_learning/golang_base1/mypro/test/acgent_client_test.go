package test

import (
	"context"
	"fmt"
	"io"
	"log"
	"myagent/protoc/pop_agent"
	"testing"
	"time"

	"google.golang.org/grpc"
)

func TestagentClient(t *testing.T) {
	conn, err := grpc.Dial(":8888", grpc.WithInsecure())
	if err != nil {
		log.Fatal("connet grpc failed: ", err)
		return
	}

	defer conn.Close()
	fmt.Println("conneted successed!!")

	//初始化客户端 并发送请求
	client := pop_agent.NewPopAgentClient(conn)

	stream, err := client.BiStream(context.Background())
	if err != nil {
		log.Fatal("BiStream failed: ", err)
		return
	}
	//收
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				fmt.Println("read done")

			}

			if err != nil {
				fmt.Println("error : ", err)

			}

			fmt.Println("recv data from server : ", in)
			fmt.Printf("%T", in)
		}
	}()

	//发
	for {
		stream.Send(&pop_agent.StreamMsg{
			CmdName: "add",
			Uuid:    "123456780",
			Content: "hello server",
			Code:    "0x00",
		})
		time.Sleep(time.Second)
	}

}
