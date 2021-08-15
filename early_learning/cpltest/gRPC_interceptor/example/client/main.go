package main

import (
	"log"
	pb "myclient/protoc/hi" // 引入proto包
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials" // 引入grpc认证包
	"google.golang.org/grpc/grpclog"
)

const (
	// Address gRPC服务地址
	Address = "127.0.0.1:9999"
)

var IsTls = true

// myCredential 自定义认证
type myCredential struct{}

// GetRequestMetadata 实现自定义认证接口
func (c myCredential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"appid":  "myappid",
		"appkey": "mykey",
	}, nil
}

// RequireTransportSecurity 自定义认证是否开启TLS
func (c myCredential) RequireTransportSecurity() bool {
	return IsTls
}

// 客户端拦截器
func Clientinterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	start := time.Now()
	err := invoker(ctx, method, req, reply, cc, opts...)
	log.Printf("method == %s ; req == %v ; rep == %v ; duration == %s ; error == %v\n", method, req, reply, time.Since(start), err)
	return err
}

func main() {
	log.SetFlags(log.Ltime | log.Llongfile)
	// TLS连接  记得把xxx改成你写的服务器地址

	var err error
	var opts []grpc.DialOption

	if IsTls {
		//打开tls 走tls认证
		creds, err := credentials.NewClientTLSFromFile("./keys/server.pem", "www.eline.com")
		if err != nil {
			log.Panicf("Failed to create TLS mycredentials %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}

	// 自定义认证，new(myCredential 的时候，由于我们实现了上述2个接口，因此new的时候，程序会执行我们实现的接口
	opts = append(opts, grpc.WithPerRPCCredentials(new(myCredential)))

	// 加上拦截器
	opts = append(opts, grpc.WithUnaryInterceptor(Clientinterceptor))

	conn, err := grpc.Dial(Address, opts...)
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

	// 故意再调用一次
	res, err = c.SayHi(context.Background(), req)
	if err != nil {
		log.Panicln(err)
	}

	log.Println(res.Message)
}
