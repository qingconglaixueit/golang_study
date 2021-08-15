package main

import (
	"github.com/tidwall/gjson"
	"log"
)

const json = `
	{"author": "xiaomotong", "age": 18, "hobby":"play"}
	{"author": "xiaozhu", "age": 19 , "hobby":"eat"}
	{"author": "zhangsan", "age": 20, "hobby":"drink"}
	{"author": "lisi", "age": 21, "hobby":"sleep"}`

func main() {

	// 设置参数，打印行数
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	// 输出 json 行数组的长度
	log.Println(gjson.Get(json, "..#"))
	// 输出 json 行 数组的第 3 行
	log.Println(gjson.Get(json, "..2"))
	// 输出 json 每一行 里面的 author 对应的值，组成一个数组
	log.Println(gjson.Get(json, "..#.author"))
	// 输出输出 json 行 中，author = xiaomotong 所在行 对应的 age 值
	log.Println(gjson.Get(json, `..#(author="xiaomotong").hobby`))

	// 遍历 json 行
	gjson.ForEachLine(json, func(jLine gjson.Result) bool {
		log.Println("author:", gjson.Get(jLine.String(), "hobby"))
		return true
	})
}
