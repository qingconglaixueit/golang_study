// 2-4-route-param.go *路由携带参数, 匹配路由
package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.RedirectTrailingSlash = false
	r.GET("/users/*id", func(c *gin.Context) {
		id := c.Param("id")
		c.String(200, "The user id is  %s", id)
	})
	// http://localhost:8080/users
	// r.GET("/users", func(c *gin.Context) {
	// 	c.String(200, "这是真正的/users")
	// })
	r.Run(":8080")
}
