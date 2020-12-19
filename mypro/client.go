package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"myagent/protoc/pop_agent"
	"time"

	"google.golang.org/grpc"
)

func StartUp() {
	err := Connect()
	if err != nil {
		fmt.Println("connect failed !!")
		return
	}
}

func Connect() error {
	connFlag := 0
connect:
	conn, err := grpc.Dial(":9999", grpc.WithInsecure())
	if err != nil {
		log.Fatal("connet grpc failed: ", err)
	}
	if connFlag == 1 {
		defer conn.Close()
	}

	fmt.Println("conneted successed!!")
	n := 1
loop:
	//初始化客户端 并发送请求
	client := pop_agent.NewPopAgentClient(conn)
	stream, err := client.BiStream(context.Background())
	if err != nil {
		fmt.Println("BiStream failed: ", err)
		time.Sleep(time.Second * time.Duration(n))
		n++
		goto loop
	}
	n = 1
	//发
	if connFlag == 1 {
		err = stream.Send(&pop_agent.StreamMsg{
			CmdName: "AgentConn",
			Uuid:    "222",
			Content: "hello server",
			Code:    "0",
		})

	} else {
		err = stream.Send(&pop_agent.StreamMsg{
			CmdName: "AgentConn",
			Uuid:    "123456780",
			Content: "hello server",
			Code:    "0",
		})
	}

	if err != nil {
		fmt.Println("send failed!!")
		goto loop
	}
	msg, err := stream.Recv()
	if err == io.EOF {
		fmt.Println("read done")
		goto loop
	}

	if err != nil {
		fmt.Println("error : ", err)
		goto loop
	}
	if err == nil {
		if msg.CmdName == "AgentConn" {
			if msg.Code == "1" {
				connFlag++
				fmt.Println("recv data from server : ", msg)
				fmt.Printf("%T", msg)
				fmt.Printf("\n你可以开始你的业务了\n")
			}
		}
	}

	if connFlag == 1 {

		fmt.Printf("\n正在处理业务...\n")
		time.Sleep(time.Second * time.Duration(1))
		fmt.Printf("\n开始请求租户信息数据\n")
		up := &pop_agent.StreamMsg{
			CmdName: "UpControl",
			Uuid:    "123456780",
			Content: "请求租户信息",
			Code:    "0",
		}
		response, err := client.DUpMsgSync(context.Background(), up)
		if err != nil {
			fmt.Println("Send error : ", err)
			return nil
		}
		fmt.Println("上报数据为：", up)
		if response.CmdName == "UpControl" {
			if response.Code == "1" {
				fmt.Println("收到服务器的响应，已经拿到租户信息", response)
				fmt.Println("本地正在处理租户信息数据，写配置文件，转租户信息数据等等...")
				fmt.Println("关闭本次连接，进行重连...")
				conn.Close()
				goto connect
			}
		}

	}
	if connFlag == 2 {
		//获取pop实时下发的数据，包括租户，运营系统的实时数据
		go func() {
			for {
				msg, err := stream.Recv()
				if err == io.EOF {
					fmt.Println("read done")
					return
				}

				if err != nil {
					fmt.Println("error : ", err)
					return
				}
				if err == nil {
					fmt.Println("----收到来自pop端的实时增量数据：", msg)

				}
			}
		}()

		for {
			fmt.Println("开始拉取全量的配置信息")
			up := &pop_agent.StreamMsg{
				CmdName: "AllConfig",
				Uuid:    "123456780",
				Content: "请求全量信息",
				Code:    "0",
			}
			response, err := client.DUpMsgSync(context.Background(), up)
			if err != nil {
				fmt.Println("AllConfig Send error : ", err)
				time.Sleep(time.Second * time.Duration(1))
				continue
				//connFlag = 1
				//go loop
			}
			if err == nil {
				fmt.Println("--收到全量的配置信息", response)

			}
			time.Sleep(time.Second * time.Duration(1))

		}
	}

	return nil
}

func main() {

	StartUp()

}
