package test

import (
	"fmt"
	"io"
	"log"
	"myagent/protoc/pop_agent"
	"net"
	"testing"

	"google.golang.org/grpc"
)

type haha struct {
}

func (h *haha) BiStream(stream pop_agent.PopAgent_BiStreamServer) error {

	//go func() {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("read done")

		}

		if err != nil {
			fmt.Println("error : ", err)

		}

		fmt.Println("recv data from client : ", in)
		fmt.Printf("%T", in)
	}
	//}()

	//发
	// for {
	// 	stream.Send(&pop_agent.StreamMsg{
	// 		CmdName: "add",
	// 		Uuid:    "99999999999",
	// 		Content: "hello client",
	// 		Code:    "0x00",
	// 	})
	// 	time.Sleep(time.Second)
	// }

}

func Testagentsrv(t *testing.T) {
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
