// 7-3-basic-auth.go
package main

import (
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(gin.BasicAuth(gin.Accounts{
		"admin": "123456",
	}))
	

	r.GET("/", func(c *gin.Context) {
		body, _ := ioutil.ReadAll(c.Request.Body)
		fmt.Println("---body--- \r\n " + string(body))
		fmt.Println("---header--- \r\n")
		for k, v := range c.Request.Header {
			fmt.Println(k, v)
		}
		fmt.Println("进入主页")
		c.JSON(200, "首页")
	})

	r.Run(":8080")
}
