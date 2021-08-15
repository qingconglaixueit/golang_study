// 8-1-json-xml-protobuf.go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
)

// 分别在浏览器中输入以下内容，会得到具体的响应信息。
// http://127.0.0.1:8080/someJSON
// http://127.0.0.1:8080/someStruct
// http://127.0.0.1:8080/someXML
// http://127.0.0.1:8080/someYAML
// http://127.0.0.1:8080/someProtoBuf

func main() {
	r := gin.Default()
	//1. json响应
	r.GET("/someJSON", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "someJSON", "status": 200})
	})
	//2. 结构体响应
	r.GET("/someStruct", func(c *gin.Context) {
		var msg struct {
			Name    string
			Message string
			Number  int
		}
		msg.Name = "root"
		msg.Message = "message"
		msg.Number = 123
		c.JSON(200, msg)
	})

	//3. XML
	r.GET("/someXML", func(c *gin.Context) {
		c.XML(200, gin.H{"message": "abc"})
	})

	//4. YAML响应
	r.GET("/someYAML", func(c *gin.Context) {
		c.YAML(200, gin.H{"name": "you"})
	})

	//5.Protobuf格式，谷歌开发的高效存储读取的工具
	r.GET("/someProtoBuf", func(c *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		//定义数据
		label := "label"
		//传protobuf格式数据
		data := &protoexample.Test{
			Label: &label,
			Reps:  reps,
		}
		c.ProtoBuf(200, data)
	})

	r.Run(":8080")
}
