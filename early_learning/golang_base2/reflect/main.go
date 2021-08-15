package main

import (
	"fmt"
	"reflect"
)

type stu struct {
	name string
	age  int
}

func modifyVarByReflect() {
	var a int = 20
	v := reflect.ValueOf(a)
	v.SetInt(30)
	fmt.Println(a)

	fmt.Println("====================")
	//
	//p := reflect.ValueOf(&a)
	//p.Elem().SetInt(10)
	//fmt.Println(a)
}
//使用反射包，修改接口的值
func main() {

	//s := &stu{"xiaozhu", 28}
	//fmt.Println(reflect.ValueOf(s))
	//fmt.Println(reflect.TypeOf(s))

	//modifyVarByReflect()

	var a int = 20
	v := reflect.ValueOf(a)
	v.SetInt(30)
	fmt.Println(a)
}
