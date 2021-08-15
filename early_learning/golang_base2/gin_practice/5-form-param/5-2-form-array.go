// 5-2-form-array.go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getting(c *gin.Context) {

}

func main() {
	// 1创建路由,默认使用了两个中间件Logger(),Recovery()
	r := gin.Default()
	// 指明html加载文件目录
	r.LoadHTMLGlob("./*")
	// 方法网页
	r.Handle("GET", "/", func(context *gin.Context) {
		fmt.Println("/访问")
		// 返回HTML文件，响应状态码200，html文件名为index.html，模板参数为nil
		context.HTML(http.StatusOK, "5-2-form-array.html", nil)
	})
	// 2绑定路由规则,
	// gin.Context,封装了request和respose
	// 表单提交处理
	r.POST("/form", func(c *gin.Context) {
		fmt.Println("/form访问")
		body, _ := ioutil.ReadAll(c.Request.Body)
		fmt.Println("---body--- \r\n " + string(body))
		fmt.Println("---header--- \r\n")
		for k, v := range c.Request.Header {
			fmt.Println(k, v)
		}
		typeStr := c.DefaultQuery("type", "alert")
		username := c.PostForm("username")
		password := c.PostForm("password")
		// 多选框
		hobbys := c.PostFormArray("hobby")
		fmt.Println("hobbys:", hobbys)
		c.String(http.StatusOK, fmt.Sprintf("type is %s, username:%s, password:%s, hooby:%v",
			typeStr, username, password, hobbys))
	})
	// 3监听端口，默认8080
	r.Run(":8080")
}
