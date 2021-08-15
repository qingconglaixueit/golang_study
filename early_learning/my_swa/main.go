package main

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
	_ "myswa/docs"
)

// gin 的处理函数  Hello
func Hello(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"msg": "hello wrold xiaomotong" })
}
// @title Xiaomotong Swagger  API
// @version 1.0
// @description 参加更文挑战第 26 天了，主题是 Swagger
// @termsOfService https://juejin.cn/user/3465271329953806

// @contact.name https://juejin.cn/user/3465271329953806
// @contact.url https://juejin.cn/user/3465271329953806
// @contact.email xxx@xxx.com.cn


// @host 127.0.0.1:8888
// @BasePath /api/v1
func main() {

	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 路由分组， 第一个版本的api  v1
	v1 := r.Group("/api/v1")
	{
		v1.GET("/hello", Hello)

	}

	// 监听端口为 8888
	r.Run(":8888")
}