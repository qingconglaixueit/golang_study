package main

import (
	"fmt"
	"pro_test2/mypro"

	"github.com/golang/protobuf/proto"
)

type Des struct {
	Name     string
	ShowName string
	Des      string
}

func main() {

	//模拟可视化包装 protobuf数据
	user1 := &mypro.User{
		Name:     "xiaozhu",
		Showname: "woshixiaozhu",
		Des:      "this is a name",
	}
	fmt.Println(user1)

	buf, _ := proto.Marshal(user1)

	inf := &mypro.Info{}
	inf.Cmd = "add"
	inf.Content = buf

	buffer, _ := proto.Marshal(inf)
	fmt.Println("序列化之后的信息为：", buffer)

	//模拟解析protobuf数据

	data := &mypro.Info{}
	proto.Unmarshal(buffer, data)
	fmt.Println("反序列化之后的信息为：", data)
	fmt.Println("cmd", data.Cmd)
	con := &mypro.User{}
	proto.Unmarshal(data.Content, con)

	fmt.Println("content", con)
	fmt.Println("name:", con.Name)
	// 	var jdata Post
	// 	err = json.Unmarshal(data.Jsondata, &jdata)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// 	fmt.Println("cmd : ", data.Cmd)
	// 	fmt.Println("jdata  name  : ", jdata.Name)
	// 	fmt.Println("jdata  showname  : ", jdata.ShowName)
	// 	fmt.Println("jdata  des  : ", jdata.Des)

	// 	//模拟包装新的 json
	// 	qqInfo := &QQ{jdata.Name, jdata.ShowName, jdata.Des}
	// 	bacInfo, err := json.MarshalIndent(qqInfo, "", "\t") //一个字段占一行

	// 	if err != nil {
	// 		fmt.Println(nil)
	// 	}
	// 	fmt.Println("QQ get  data : \n", string(bacInfo))
}
