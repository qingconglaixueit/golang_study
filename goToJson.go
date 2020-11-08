/*
	author:qingbing
	date: 2020年11月 8日 20:01:10
	what: go数据对象 -> json数据
*/
package main

import (
	"encoding/json" //json包
	"fmt"
)

//将结构体转成 json ,结构体成员变量第一个字母必须大写，只有字段首字母大写的才会被转换
type Post struct {
	Name     string
	ShowName string
	Des      string
}

type Stu struct {
	Info   map[string]string `json:"个人信息"`
	Hobby  []string          `json:"爱好"`
	Enable bool              `json:"是否启用"`
	test   []string          `json:"tttt"`
}

func main() {

	//将结构体转成 json
	post := &Post{"xiaozhu", "hello xiaozhu", "this is a name"}
	b, err := json.MarshalIndent(post, "", "\t") //一个字段占一行

	if err != nil {
		fmt.Println(nil)
	}
	fmt.Println(string(b))

	//slice / array 转换成json 数据
	post2 := []string{"apple", "pen", "penapple"}
	b2, err2 := json.Marshal(post2)
	if err2 != nil {
		fmt.Println(nil)
	}
	fmt.Println(string(b2))

	//集合map转换成 json 数据 ， map 的 key 必须是string类型，这是json要求的
	post3 := map[string]string{"name": "xiaozhu", "age": "13", "height": "168"}
	b3, err3 := json.MarshalIndent(post3, "", "\t")
	if err3 != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b3))

	//组合在一起使用
	var stu Stu
	stu.Enable = true
	stu.Hobby = []string{"running", "basketball", "football"}
	stu.Info = map[string]string{
		"name":   "xiaozhu",
		"age":    "24",
		"height": "178"}
	bs, errs := json.MarshalIndent(&stu, "", "\t")
	//bs, errs := json.Marshal(&stu)
	if errs != nil {
		fmt.Println(errs)
	}
	fmt.Println(string(bs))
}
