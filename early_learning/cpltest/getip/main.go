package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lionsoul2014/ip2region/binding/golang/ip2region"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
		fmt.Println("客户端IP：", c.ClientIP())
		region, err := ip2region.New("./ip2region.db ")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(region.BinarySearch(c.ClientIP()))
	})
	r.Run("10001") // listen and serve on 0.0.0.0:8080
}
