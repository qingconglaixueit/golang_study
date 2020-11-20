package main

import (
	"encoding/json"
	"fmt"
	"pro_test1/mypro"

	"github.com/golang/protobuf/proto"
)

//将结构体转成 json ,结构体成员变量第一个字母必须大写，只有字段首字母大写的才会被转换
type Post struct {
	Name     string
	ShowName string
	Des      string
}

type QQ struct {
	Name     string
	ShowName string
	Des      string
}

func main() {

	//模拟可视化包装 protobuf数据
	post := &Post{"xiaozhu", "hello xiaozhu", "this is a name"}
	b, err := json.MarshalIndent(post, "", "\t") //一个字段占一行

	if err != nil {
		fmt.Println(nil)
	}
	fmt.Println(string(b))

	info := &mypro.Info{}
	info.Jsondata = b
	info.Cmd = "AddUser"

	buffer, _ := proto.Marshal(info)
	fmt.Println("序列化之后的信息为：", buffer)

	//模拟解析protobuf数据

	data := &mypro.Info{}
	proto.Unmarshal(buffer, data)
	fmt.Println("反序列化之后的信息为：", data)

	var jdata Post
	err = json.Unmarshal(data.Jsondata, &jdata)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("cmd : ", data.Cmd)
	fmt.Println("jdata  name  : ", jdata.Name)
	fmt.Println("jdata  showname  : ", jdata.ShowName)
	fmt.Println("jdata  des  : ", jdata.Des)

	//模拟包装新的 json
	qqInfo := &QQ{jdata.Name, jdata.ShowName, jdata.Des}
	bacInfo, err := json.MarshalIndent(qqInfo, "", "\t") //一个字段占一行

	if err != nil {
		fmt.Println(nil)
	}
	fmt.Println("QQ get  data : \n", string(bacInfo))
}
