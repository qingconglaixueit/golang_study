// 7-4-global.go 全局中间件
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
	c.Next() //调用后续的处理函数 执行indexHandler函数
	//c.Abort() //阻止调用后续的处理函数
	cost := time.Since(start)
	fmt.Println("cost:%v\n", cost)
	fmt.Println("m1 out")
}
func m2(c *gin.Context) {
	fmt.Println("m2 in ....")
	//c.Next() //调用后续的处理函数
	c.Abort() //阻止后续调用
	//return   //return 立即结束m2函数
	//m1 in ....
	//m2 in ....
	//cost:%v
	// 0s
	//m1 out
	fmt.Println("m2 out")
}

//func authMiddleware(c *gin.Context)  {   //通常写成闭包
//	//是否登陆的判断
//	//if 是登陆用户
//	//c.Next()
//	//else
//	//c.Abort()
//}
func authMiddleware(doCheck bool) gin.HandlerFunc { //开关注册
	//连接数据库
	//或着其他准备工作
	return func(c *gin.Context) {
		if doCheck {
			//是否登陆的判断
			//if 是登陆用户
			//c.Next()
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
	r.GET("/shop", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "index",
		})
	})
	r.GET("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "index",
		})
	})
	r.Run(":8080")
}
