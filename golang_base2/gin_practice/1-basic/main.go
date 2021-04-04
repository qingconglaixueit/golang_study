// 1-basic
package main

import "github.com/gin-gonic/gin"

func main() {
	// Default方法的主要作用是实例化一个带有日志、故障恢复中间件的引擎。
	r := gin.Default() //实例化一个gin对象
	r.RedirectTrailingSlash = false //进制重定向
	// 定义请求
	//定义一个GET请求的路由，参数一是路由地址，也就是在浏览器访问的相对路径，
	// 					    参数二是一个匿名函数，函数内部用于业务逻辑处理。
	r.GET("/login", func(c *gin.Context) {
		//tmp := c.Param("id")
		c.JSON(200, gin.H{ //JSON内容可以通过gin提供的H方法来构建，非常方便。
			"msg": "login", //调用JSON方法返回数据。JSON的操作非常简单，参数一是状态码，参数二是JSON的内容。
		})
		// c.XML()
		// c.HTML()
	})

	//r.GET("/login/*id", func(c *gin.Context) {
	//	tmp := c.Param("id")
	//	c.JSON(200, gin.H{ //JSON内容可以通过gin提供的H方法来构建，非常方便。
	//		"msg": tmp, //调用JSON方法返回数据。JSON的操作非常简单，参数一是状态码，参数二是JSON的内容。
	//	})
	//	// c.XML()
	//	// c.HTML()
	//})
	// Run方法最终会调用内置http库的ListenAndServe方法来监听端口，如果不传参数默认监听80端口，
	// 也可以通过参数来变更地址和端口。
	r.Run(":12005")
}
