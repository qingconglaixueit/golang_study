// 7-1-basic.go

package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func indexHandler(c *gin.Context) {
	fmt.Println("index in ...")
	c.JSON(http.StatusOK, gin.H{
		"msg": "indx",
	})
}

func costTime() gin.HandlerFunc {
	return func(c *gin.Context) {
		//请求前获取当前时间
		nowTime := time.Now()
		fmt.Println("获取起始时间")
		//请求处理
		c.Next() // 处理handler
		//处理后获取消耗时间
		fmt.Println("获取结束时间")
		costTime := time.Since(nowTime)
		url := c.Request.URL.String()
		fmt.Printf("the request URL %s cost %v\n", url, costTime)
	}
}

//定义一个中间件
func m1(c *gin.Context) {
	fmt.Println("m1 in ....")
}

func m2(c *gin.Context) {
	fmt.Println("m2 in ....")
}
func main() {
	r := gin.Default()
	r.Use(costTime())
	//GET(relativePath string, handlers ...HandlerFunc) IRoutes
	r.GET("/index", m1, m2, indexHandler)
	r.GET("/", func(c *gin.Context) {
		fmt.Println("处理过程")
		c.JSON(200, "首页")
		// c.Next()
	})
	r.Run(":8080")
}
