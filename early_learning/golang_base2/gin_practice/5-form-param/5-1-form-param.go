// 5-1-form-param.go
package main

import (
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

// curl -d "message=pingye" http://localhost:8080/form_post
// {"message":"pingye","nick":"anonymous","status":"posted"}
func main() {
	r := gin.Default()
	r.POST("/form_post", func(c *gin.Context) {
		body, _ := ioutil.ReadAll(c.Request.Body)
		fmt.Println("---body--- \r\n " + string(body))
		//fmt.Println("---header--- \r\n")
		//for k, v := range c.Request.Header {
		//	fmt.Println(k, v)
		//}
		message := c.PostForm("message")
		nick := c.PostForm("nick")

		c.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})
	r.Run(":8080")
}
