// 6-1-upload-file.go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// 1创建路由,默认使用了两个中间件Logger(),Recovery()
	r := gin.Default()
	// 给表单限制上传大小 (默认 32 MiB)
	r.MaxMultipartMemory = 8 << 20 // 8 MiB
	r.Static("/", "./test")
	// 2绑定路由规则,
	// gin.Context,封装了request和respose
	r.POST("/upload", func(c *gin.Context) {
		

		file, _ := c.FormFile("file")
		log.Println("file:", file.Filename)
		c.SaveUploadedFile(file, "./"+"test/"+file.Filename) // 上传文件到指定的路径
		c.String(200, fmt.Sprintf("%s upload file!", file.Filename))
	})
	// 3监听端口，默认8080
	r.Run(":8080")
}
