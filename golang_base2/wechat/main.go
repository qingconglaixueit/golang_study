package main

import (
	"fmt"
	"github.com/wonderivan/logger"
	"net/http"
	"time"
)

const (
	logLevel = "dev"
	port     = 80
	token    = "YrD1PLK6rxPTKCoW1ZKM2SS4zCGCG5OR"
)

// 编译方法
// go mod init wechat
// go build
// ./wechat
// 需要自己修改token，以适应自己公众号的token
func main() {

	logger.SetLogger("./log.json")
	logger.Info(" ------------ start  main ------------")

	server := http.Server{
		Addr:           fmt.Sprintf(":%d", port), // 设置监听地址， ip:port
		Handler:        &httpHandler{},           // 用什么handler来处理
		ReadTimeout:    5 * time.Second,          // 读写超时 微信给出来5
		WriteTimeout:   5 * time.Second,
		MaxHeaderBytes: 0,
	}

	logger.Info(fmt.Sprintf("Listen: %d", port))
	logger.Fatal(server.ListenAndServe())
}
