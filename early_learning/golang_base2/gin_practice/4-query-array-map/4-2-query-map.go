// 4-1-query-array.go
package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// 现在运行代码，访问http://localhost:8080/map?ids[a]=123&ids[b]=456&ids[c]=789，就会看到如下信息：
// {"a":"123","b":"456","c":"789"}

func main() {

	r := gin.Default()
	r.GET("/map", func(c *gin.Context) {
		fmt.Println("map:", c.QueryMap("ids"))
		c.JSON(200, c.QueryMap("ids"))
	})
	r.GET("/", func(c *gin.Context) {
		fmt.Println("map:", c.QueryMap("ids"))
		c.JSON(200, c.QueryMap("ids"))
	})
	r.Run(":8080")
}
