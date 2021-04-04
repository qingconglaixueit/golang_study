// 7-5-self-meddle.go  针对特定的url授权
package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func costTime() gin.HandlerFunc {
	return func(c *gin.Context) {
		//请求前获取当前时间
		nowTime := time.Now()
		fmt.Println("获取起始时间")
		//请求处理
		// c.Next() // 处理handler
		//处理后获取消耗时间
		fmt.Println("获取结束时间")
		costTime := time.Since(nowTime)
		url := c.Request.URL.String()
		fmt.Printf("the request URL %s cost %v\n", url, costTime)
	}
}
func main() {
	r := gin.New()
	r.Use(costTime())
	// 统计每个处理的耗时
	r.GET("/", func(c *gin.Context) {
		fmt.Println("处理过程")
		c.JSON(200, "首页")
		// c.Next()
	})
	r.Run(":8080")
}
