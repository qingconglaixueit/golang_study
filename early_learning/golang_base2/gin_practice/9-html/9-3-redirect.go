// 9-3-redirect.go

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("view/*")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"title": "我是gin", "name": "you"})
	})
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/index")	// 重定向
	})
	r.Run(":8080")
}
