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

//定义一个中间件:统计耗时
func m1(c *gin.Context) {
	fmt.Println("m1 in ....")
	//计时
	start := time.Now()
	// c.Abort() //阻止调用后续的处理函数
	c.Next() //调用后续的处理函数 执行indexHandler函数

	cost := time.Since(start)
	fmt.Println("cost:%v\n", cost)
	//输出
	// m1 in ....
	//index in ...
	//cost:%v
	// 996.8µs
}
func main() {
	r := gin.Default()
	//GET(relativePath string, handlers ...HandlerFunc) IRoutes
	r.GET("/index", m1, indexHandler) //先执行m1函数再执行indexHandler函数

	r.Run(":8080")
}
