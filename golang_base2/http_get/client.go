package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//client端
func main() {
	//get请求
	resp, err := http.Get("127.0.0.1:8888")
	if err != nil {
		fmt.Printf("Get err : %s", err)
		return
	}
	//defer 关闭连接
	defer resp.Body.Close()

	//body是[]byte流
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("ReadAll err : %s", err)
		return
	}

	fmt.Println(string(body))


	fmt.Println(resp.Status)
	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Proto)
	fmt.Println(resp.ProtoMajor)
	fmt.Println(resp.ProtoMinor)

}
