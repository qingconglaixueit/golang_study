// 10-1-sync-async.go
package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// http://127.0.0.1:8080/long_async
// http://127.0.0.1:8080/long_sync
func main() {
	r := gin.Default()
	//1. 异步
	r.GET("/long_async", func(c *gin.Context) {
		//需要搞一个副本
		copyContext := c.Copy()
		//异步处理
		go func() {
			time.Sleep(3 * time.Second)
			log.Println("异步执行：" + copyContext.Request.URL.Path)
			// copyContext.JSON(200, gin.H{"message": "someJSON", "status": 200})
		}()
	})

	//2. 同步
	r.GET("/long_sync", func(c *gin.Context) {
		time.Sleep(3 * time.Second)
		log.Println("同步执行：" + c.Request.URL.Path)
	})
	r.Run()
}
