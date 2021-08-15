// 3-2-url-param.go url参数获取
package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.DefaultQuery("id", "0")
		value, ok := c.GetQuery("id") // 适合用来判断是否存在该参数

		if ok {
			fmt.Println("id:", value)
		} else {
			fmt.Println("id: nil")
		}

		c.String(200, c.DefaultQuery("wechat", "default baidu_org"))
	})
	r.Run(":8080")
}
