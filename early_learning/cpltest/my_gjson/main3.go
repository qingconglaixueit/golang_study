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

	// 获取名字
	log.Println("author:", gjson.Get(json, "author.name"))
	// 获取年龄
	log.Println("age:", gjson.Get(json, "age"))

	log.Println("hobby:", gjson.Get(json, "hobb?"))
	log.Println("hobby count:", gjson.Get(json, "hobby.#"))

	log.Println("second hobby:", gjson.Get(json, "ho?by.1"))
	log.Println("third hobby:", gjson.Get(json, "ho*.2"))

	log.Println("love.music", gjson.Get(json, `love\.music`))

	log.Println("location first city :", gjson.Get(json, "location.0"))
	log.Println("location second city :", gjson.Get(json, "location.1"))
}
