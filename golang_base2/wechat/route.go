package main

import (
	"fmt"
	"github.com/wonderivan/logger"
	"io"
	"net/http"
	"regexp"
	"time"
	"wechat/wx"
)

type WebController struct {
	Function func(http.ResponseWriter, *http.Request)
	Method   string
	Pattern  string
}

var mux []WebController // 自己定义的路由
// ^ 匹配输入字符串的开始位置
func init() {
	mux = append(mux, WebController{post, "POST", "^/"})
	mux = append(mux, WebController{get, "GET", "^/"})
}

type httpHandler struct{} // 实际是实现了Handler interface


func (*httpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	t := time.Now()
	for _, webController := range mux { // 遍历路由
		// 匹配请求的   r.URL.Path  -> webController.Pattern
		if m, _ := regexp.MatchString(webController.Pattern, r.URL.Path); m { // 匹配URL
			logger.Info("webController.Pattern  == ", webController.Pattern)
			logger.Info("r.URL.Path  == ", r.URL.Path)
			if r.Method == webController.Method { // 匹配方法
				logger.Info("webController.Method  == ", webController.Method)

				webController.Function(w, r) // 调用对应的处理函数

				d := time.Now().Sub(t)

				l := fmt.Sprintf("[ACCESS] | % -10s | % -40s | % -16s", r.Method, r.URL.Path, d.String())

				logger.Info(l)

				return
			}
		}
	}

	d := time.Now().Sub(t)

	l := fmt.Sprintf("[ACCESS] | % -10s | % -40s | % -16s", r.Method, r.URL.Path, d.String())

	logger.Info(l)

	io.WriteString(w, "")
	return
}

// 处理token的认证
func get(w http.ResponseWriter, r *http.Request) {

	client, err := wx.NewClient(r, w, token)

	if err != nil {
		logger.Info(err)
		w.WriteHeader(403) // 校验失败
		return
	}

	if len(client.Query.Echostr) > 0 {
		logger.Info("Echostr == ", client.Query.Echostr)
		w.Write([]byte(client.Query.Echostr)) // 校验成功返回的是Echostr
		return
	}

	w.WriteHeader(403)
	return
}

// 微信平台过来消息， 处理 ，然后返回微信平台
func post(w http.ResponseWriter, r *http.Request) {

	client, err := wx.NewClient(r, w, token)

	if err != nil {
		logger.Info(err)
		w.WriteHeader(403)
		return
	}
	// 到这一步签名已经验证通过了
	client.Run()
	return
}
