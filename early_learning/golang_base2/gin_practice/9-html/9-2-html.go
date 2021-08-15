// 9-2-html.go 复杂一些的例子
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("view2/**/*")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "user/index.html", gin.H{"title": "我是gin", "name": "you2"})
	})
	r.Run()
}
