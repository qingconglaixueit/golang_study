// 2-5-route-param.go *路由携带参数,禁止重定向
package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.RedirectTrailingSlash = false // 禁止重定向
	r.GET("/users/*id", func(c *gin.Context) {
		id := c.Param("id")
		c.String(200, "The user id is  %s", id)
	})
	r.Run(":8080")
}
