// 7-1-route-group.go
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//V1版本的API
	v1Group := r.Group("/v1")
	v1Group.GET("/users", func(c *gin.Context) {
		c.String(200, "/v1/users")
	})
	v1Group.GET("/products", func(c *gin.Context) {
		c.String(200, "/v1/products")
	})
	//V2版本的API
	v2Group := r.Group("/v2")
	v2Group.GET("/users", func(c *gin.Context) {
		c.String(200, "/v2/users")
	})
	v2Group.GET("/products", func(c *gin.Context) {
		c.String(200, "/v2/products")
	})

	//V3版本的API
	v3Group := r.Group("/v3")
	v3Group.GET("/users", func(c *gin.Context) {
		c.String(200, "/v3/users")
	})
	v3Group.GET("/products", func(c *gin.Context) {
		c.String(200, "/v3/products")
	})
	r.Run(":8080")
}
