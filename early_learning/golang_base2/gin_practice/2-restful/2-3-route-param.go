// 2-3-route-param.go *路由携带参数
package main

import "github.com/gin-gonic/gin"

// 浏览器里访问http://localhost:8080/users/123，会看到如下信息：
// The user id is  /123
// 我们获取到的id不是123了，而是/123，多了一个/
func main() {
	r := gin.Default()
	r.GET("/users/*id", func(c *gin.Context) {
		id := c.Param("id")
		c.String(200, "The user id is  %s", id)
	})
	r.Run(":8080")
}
