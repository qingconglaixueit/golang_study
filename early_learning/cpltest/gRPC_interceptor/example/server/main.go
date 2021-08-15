package main

import (
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"log"
	"net"

	pb "myserver/protoc/hi"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials" // 引入grpc认证包
)

const (
	// Address gRPC服务地址
	Address = "127.0.0.1:9999"
)

// 定义helloService并实现约定的接口
type HiService struct{}

// HiService Hello服务
var HiSer = HiService{}

// SayHello 实现Hello服务接口
func (h HiService) SayHi(ctx context.Context, in *pb.HiRequest) (*pb.HiResponse, error) {

	// 解析metada中的信息并验证
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, grpc.Errorf(codes.Unauthenticated, "no token ")
	}

	var (
		appId  string
		appKey string
	)

	// md 是一个 map[string][]string 类型的
	if val, ok := md["appid"]; ok {
		appId = val[0]
	}

	if val, ok := md["appkey"]; ok {
		appKey = val[0]
	}

	if appId != "myappid" || appKey != "mykey" {
		return nil, grpc.Errorf(codes.Unauthenticated, "token invalide: appid=%s, appkey=%s", appId, appKey)
	}

	resp := new(pb.HiResponse)
	resp.Message = fmt.Sprintf("Hi %s.", in.Name)

	return resp, nil
}

// 认证token
func myAuth(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return grpc.Errorf(codes.Unauthenticated, "no token ")
	}

	log.Println("myAuth ...")

	var (
		appId  string
		appKey string
	)

	// md 是一个 map[string][]string 类型的
	if val, ok := md["appid"]; ok {
		appId = val[0]
	}

	if val, ok := md["appkey"]; ok {
		appKey = val[0]
	}

	if appId != "myappid" || appKey != "mykey" {
		return grpc.Errorf(codes.Unauthenticated, "token invalide: appid=%s, appkey=%s", appId, appKey)
	}

	return nil
}

// interceptor 拦截器
func interceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// 进行认证
	log.Println("interceptor...")
	err := myAuth(ctx)
	if err != nil {
		return nil, err
	}

	// 继续处理请求
	return handler(ctx, req)
}

func main() {
	log.SetFlags(log.Ltime | log.Llongfile)

	listen, err := net.Listen("tcp", Address)
	if err != nil {
		log.Panicf("Failed to listen: %v", err)
	}

	var opts []grpc.ServerOption

	// TLS认证
	creds, err := credentials.NewServerTLSFromFile("./keys/server.pem", "./keys/server.key")
	if err != nil {
		log.Panicf("Failed to generate credentials %v", err)
	}

	opts = append(opts, grpc.Creds(creds))

	// 注册一个拦截器
	opts = append(opts, grpc.UnaryInterceptor(interceptor))
	// 实例化grpc Server, 并开启TLS认证，其中还有拦截器
	s := grpc.NewServer(opts...)

	// 注册HelloService
	pb.RegisterHiServer(s, HiSer)

	log.Println("Listen on " + Address + " with TLS and interceptor")

	s.Serve(listen)
}
