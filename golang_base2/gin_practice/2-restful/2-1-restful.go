// 2-1-restful.go  基本restful api
package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// 测试方法
// curl -X PUT http://localhost:8080/article
// curl -X POST http://localhost:8080/article
// curl -X GET http://localhost:8080/article
// curl -X DELETE http://localhost:8080/article

func main() {
	router := gin.Default()
	// 请求动词的第一个参数是请求路径，第二个参数是用于逻辑处理的函数
	router.POST("/article", func(c *gin.Context) {
		c.String(200, "article post")
	})
	router.DELETE("/article", func(c *gin.Context) {
		c.String(200, "article delete")
	})
	router.PUT("/article", func(c *gin.Context) {
		c.String(200, "article put")
	})
	router.GET("/article", func(c *gin.Context) {
		c.String(200, "article get")
	})
	// http://localhost:8080/article/go
	router.GET("/article/:id", func(c *gin.Context) {
		id := c.Param("id")
		fmt.Println("/article/:id ->", id)
		c.String(200, id)
	})
	// http://localhost:8080/article/go/linux
	// router.GET("/article/:id/:action", func(c *gin.Context)
	// router.GET("/article/:id/*action", func(c *gin.Context) {
	// 	id := c.Param("id")
	// 	action := c.Param("action")
	// 	fmt.Printf("1 /article/:id->%s, action:%s\n", id, action)
	// 	c.String(200, id+" "+action)
	// })

	// http://localhost:8080/article/go/linux
	router.GET("/article/:id/:action", func(c *gin.Context) {
		id := c.Param("id")
		action := c.Param("action")
		fmt.Printf("2 /article/:id->%s, action:%s\n", id, action)
		c.String(200, id+" "+action)
	})

	router.Run(":8080")
}
