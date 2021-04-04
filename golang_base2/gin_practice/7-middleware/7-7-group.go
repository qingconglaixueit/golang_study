// 7-7-group.go
package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func authMiddleware(doCheck bool) gin.HandlerFunc { //开关注册
	//连接数据库
	//或着其他准备工作
	return func(c *gin.Context) {
		if doCheck {
			//是否登陆的判断
			//if 是登陆用户
			//c.Next()
			//else
			//c.Abort()
		} else {
			c.Next()
		}
	}

}

func main() {
	r := gin.Default()
	//路由组注册中间件方法1：
	xx1Group := r.Group("/xx1", func(c *gin.Context) { fmt.Println("/v1中间件") })
	{
		xx1Group.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "xx1Group"})
		})
		xx1Group.GET("/index2", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "2222xx1Group"})
		})
	}
	//路由组注册中间件方法2：
	xx2Group := r.Group("/xx2")
	xx2Group.Use(authMiddleware(true))
	{
		xx2Group.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "xx2Group"})
		})
	}
	r.Run(":8080")
}
