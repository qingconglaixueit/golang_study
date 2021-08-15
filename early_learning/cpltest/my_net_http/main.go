package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/Hi", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>Hi xiaomotong</h1> "))
	})

	if err := http.ListenAndServe(":8888", nil); err != nil {
		fmt.Println("http server error:", err)
	}

	//route := gin.Default()
	//route.GET("/hi", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"msg": "欢迎关注小魔童哪吒",
	//	})
	//})
	//// 监听的地址和端口，默认地址是127.0.0.1
	//// 此处指定端口8888
	//route.Run(":8888")
}
