package main

import (
	"github.com/tidwall/gjson"
	"log"
)

const json = `
{
  "author":{"name":"xiaomotong", "nick": "xiaozhu"},
  "age": 18,
  "hobby": ["play", "eat", "drink"],
  "love.music": "one day",
  "location": [
    {"province": "gd", "city":"gz", "area": "huangpu"},
    {"province": "gd", "city":"sz", "area": "nanshan"},
  ]
}
`

func main() {

	// 设置参数，打印行数
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	// 翻转 hobby 数组
	log.Println("hobby reverse:", gjson.Get(json, "hobby|@reverse"))
	// 移除空白符
	log.Println("location.0:", gjson.Get(json, "location.0|@ugly"))

	// 使json 更加容易阅读 pretty
	log.Println("location : ", gjson.Get(json, "location.1|@pretty"))
	// 输出整个json
	log.Println("this : ", gjson.Get(json, "@this"))

	test := `["小猪1", ["小猪2", "小猪3"]]`
	// 扁平化
	log.Println("this : ", gjson.Get(test, "@flatten"))
}
