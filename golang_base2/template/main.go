package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var myTemp *template.Template

type Person struct {
	Name  string
	Title string
	Age   int
}

//开始处理数据和做响应
func userInfo(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	//var pp []Person
	p1 := Person{
		Name:  "小猪1号",
		Title: "看门的",
		Age:   18,
	}
	//p2 := Person{
	//	Name:  "小猪2号",
	//	Title: "洗门的",
	//	Age:   15,
	//}
	//p3 := Person{
	//	Name:  "小猪3号",
	//	Title: "做门的",
	//	Age:   29,
	//}

	//pp = append(pp, p1)

	err := myTemp.Execute(w, p1)
	if err != nil {
		fmt.Println("Execute err ；%v", err)
		return
	}

}

func main() {

	//初始化模板
	var err error
	myTemp, err = template.ParseFiles("./index.html")
	if err != nil {
		fmt.Println("ParseFiles err ；%v", err)
		return
	}

	//注册处理模板的函数  并 开启监听
	http.HandleFunc("/", userInfo)
	err = http.ListenAndServe("0.0.0.0:8888", nil)
	if err != nil {
		fmt.Println("ListenAndServe err ；%v", err)
		return
	}

}
