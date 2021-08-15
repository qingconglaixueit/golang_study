package main

import (
	"log"

	"github.com/tidwall/gjson"
)

func main() {

	// 设置参数，打印行数
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	json := `
			{
				"author": {
					"name": "xiaomotong",
					"age": 18,
					"hobby": "writing"
				},
				"extra": "hello wolrd",
				"picList":[{"name":"xiaozhu1"},{"name":"xiaozhu2"}]
			}
			`
	// 校验 json 字符串是否合法
	// 如果不合法的话， gjson 不会报错 panic，可能会拿到一个奇怪值
	if gjson.Valid(json){
		log.Println("json valid ...")
	}else{
		log.Fatal("json invalid ... ")
	}

	// 获取 author.name 的 值
	aName := gjson.Get(json, "author.name")
	log.Println("aName :", aName.String())

	// 获取 extra 的值
	extra := gjson.Get(json, "extra")
	log.Println("extra:", extra)

	// 获取 一个不存在的 键 对应的 值
	non := gjson.Get(json, "non")
	log.Println("non:", non)


	// 一次性 获取json 的多个键 值
	res := gjson.GetMany(json, "author.age", "author.hobby","picList")
	for i, v := range res{
		if i == 0{
			log.Println(v.Int())
		}else if i == 2{
			for _,vv := range v.Array(){
				log.Println("picList.name :",vv.Get("name"))
			}
		}else{
			log.Println(v)
		}
	}
}
