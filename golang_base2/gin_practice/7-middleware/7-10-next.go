package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()

	mid1 := func(c *gin.Context) {
		fmt.Println("mid1 start")
		c.Next()

		fmt.Println("mid1 end")
	}
	mid2 := func(c *gin.Context) {
		fmt.Println("mid2 start")
		//c.Next()
		c.Abort()
		fmt.Println("mid2 end")
	}
	mid3 := func(c *gin.Context) {
		fmt.Println("mid3 start")
		//c.Next()
		fmt.Println("mid3 end")
	}
	router.Use(mid1, mid2)
	router.Use(mid3)
	router.GET("/index", func(c *gin.Context) {
		fmt.Println("process get request")
		c.JSON(http.StatusOK, "hello")
		fmt.Println("JSON after") //
		// c.Next() // 这里加是没有用
	})

	router.Run(":8080")
}
