// 7-4-url-auth.go  针对特定的url授权
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, "首页")
	})
	adminGroup := r.Group("/admin")
	adminGroup.Use(gin.BasicAuth(gin.Accounts{ // 可以专门针对某一个路径去做授权
		"admin": "123456",
	}))
	adminGroup.GET("/index", func(c *gin.Context) {
		c.JSON(200, "后台首页")
	})
	r.Run(":8080")
}
