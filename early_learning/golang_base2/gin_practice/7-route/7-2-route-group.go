// 7-2-route-group.go
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	v1Group := r.Group("/v1")
	{ // 代码可读性
		v1Group.GET("/users", func(c *gin.Context) {
			c.String(200, "/v1/users")
		})
		v1Group.GET("/products", func(c *gin.Context) {
			c.String(200, "/v1/products")
		})
	}
	v1v1Group := v1Group.Group("v1") // 分组嵌套
	{                                // 代码可读性
		v1v1Group.GET("/users", func(c *gin.Context) {
			c.String(200, "/v1/v1/users")
		})
		v1v1Group.GET("/products", func(c *gin.Context) {
			c.String(200, "/v1/v1/products")
		})
	}

	v2Group := r.Group("/v2")
	{
		v2Group.GET("/users", func(c *gin.Context) {
			c.String(200, "/v2/users")
		})
		v2Group.GET("/products", func(c *gin.Context) {
			c.String(200, "/v2/products")
		})
	}
	r.Run(":8080")
}
