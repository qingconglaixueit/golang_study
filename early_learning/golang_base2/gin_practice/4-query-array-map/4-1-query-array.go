// 4-2-query-map.go
package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// 在浏览器里访问http://localhost:8080/?media=blog&media=wechat 会看到如下信息：
// ["blog","wechat"]
func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		fmt.Println("media:", c.QueryArray("media"))
		c.JSON(200, c.QueryArray("media"))
	})
	r.Run(":8080")
}
