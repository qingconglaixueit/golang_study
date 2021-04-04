package main

import (
	"github.com/gin-gonic/gin"
	"my/controller"
	"my/myauth"
)

func main() {
	//连接数据库
	conErr := controller.InitMySQLCon()
	if conErr != nil {
		panic(conErr)
	}

	//需要使用到gorm，因此需要先做一个初始化
	controller.InitModel()
	defer controller.DB.Close()


	route := gin.Default()

	//路由分组
	v1 := route.Group("/v1/")
	{
		//登录（为了方便，将注册和登录功能写在了一起）
		v1.POST("/login", controller.Login)
	}


	v2 := route.Group("/v1/auth/")
	//一个身份验证的中间件
	v2.Use(myauth.JWTAuth())
	{
		//带着token请求服务器
		v2.POST("/hello", controller.Hello)
	}

	//监听9999端口
	route.Run(":9999")

}
