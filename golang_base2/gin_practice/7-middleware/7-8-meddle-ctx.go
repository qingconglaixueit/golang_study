// 7-8-meddle-ctx.go 跨中间件取值
package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func indexHandler(c *gin.Context) {
	fmt.Println("index in ...")
	name, ok := c.Get("name") //从上下文中取值，跨中间件存取值
	if !ok {
		name = "匿名用户"
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": name,
	})
}

//定义一个中间件:统计耗时
func m1(c *gin.Context) {
	fmt.Println("m1 in ....")
	//计时
	start := time.Now()
	c.Next() //调用后续的处理函数 执行indexHandler函数
	//c.Abort() //阻止调用后续的处理函数
	cost := time.Since(start)
	fmt.Println("cost:%v\n", cost)
	fmt.Println("m1 out")
}
func m2(c *gin.Context) {
	fmt.Println("m2 in ....")
	c.Set("name", "0voice") //在上下文中设置c的值
	fmt.Println("m2 out")
}
func authMiddleware(doCheck bool) gin.HandlerFunc { //开关注册
	//连接数据库
	//或着其他准备工作
	return func(c *gin.Context) {
		if doCheck {
			//是否登陆的判断
			//if 是登陆用户
			c.Next()
			//else
			//c.Abort()
		} else {
			c.Next()
		}
	}
}
func main() {
	r := gin.Default()
	r.Use(m1, m2, authMiddleware(true)) //全局注册中间件函数m1,m2    洋葱模型   类似递归调用
	//GET(relativePath string, handlers ...HandlerFunc) IRoutes
	//r.GET("/index",m1,indexHandler)  //先执行m1函数再执行indexHandler函数
	r.GET("/index", indexHandler)
	r.Run(":8080")
}
