// 3-1-url-param.go url参数获取
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {

		//test := c.Request.URL.Query().Get("wechat")
		//fmt.Printf("!!!!!!wechat is %s\n\n\n\n", test)

		if c.Query("wechat") == "" || len(c.Query("wechat")) == 0 {
			fmt.Printf("wechat is empty\n")
		} else {
			fmt.Printf("wechat is %s\n", c.Query("wechat"))
		}
		//c.String(200, c.Query("wechat"))
		// fmt.Println("do it ?")
	})
	r.Run(":8080")


}
