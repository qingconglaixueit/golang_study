package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.Use(gin.BasicAuth(gin.Accounts{
		"admin": "123456",
	}))

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, "首页")
	})

	r.Run(":8080")
}
