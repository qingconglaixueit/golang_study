package main

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	"study.com/mypro" //引入的proto的包 ， go语言中，一个目录里面的所有文件都应设置成一个包
)

func main() {
	s1 := &mypro.Student{} //第一个学生信息
	s1.Name = "jz01"
	s1.Age = 23
	s1.Address = "cq"
	s1.Cn = mypro.ClassName_class2 //枚举类型赋值
	ss := &mypro.Students{}
	ss.Person = append(ss.Person, s1) //将第一个学生信息添加到Students对应的切片中
	s2 := &mypro.Student{}            //第二个学生信息
	s2.Name = "jz02"
	s2.Age = 25
	s2.Address = "cd"
	s2.Cn = mypro.ClassName_class3
	ss.Person = append(ss.Person, s2) //将第二个学生信息添加到Students对应的切片中
	ss.School = "cqu"
	fmt.Println("Students信息为：", ss)
	// Marshal takes a protocol buffer message
	// and encodes it into the wire format, returning the data.
	buffer, _ := proto.Marshal(ss)
	fmt.Println("序列化之后的信息为：", buffer)
	//  Use UnmarshalMerge to preserve and append to existing data.
	data := &mypro.Students{}
	proto.Unmarshal(buffer, data)
	fmt.Println("反序列化之后的信息为：", data)
}
