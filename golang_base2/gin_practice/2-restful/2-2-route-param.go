// 2-2-route-param.go :路由携带参数
package main

import "github.com/gin-gonic/gin"

// 测试 打开浏览器，输入http://localhost:8080/users/123，就可以看到如下信息：
// The user id is  123
func main() {
	r := gin.Default()
	r.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.String(200, "The user id is  %s", id)
	})
	r.Run(":8080")
}
