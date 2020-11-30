package main

import (
	"fmt"
	"reflect"
)

type ControlMapType map[string]reflect.Value

type XXAgent struct {
}

func (this *XXAgent) XXConvertUser(pro string) string {
	fmt.Printf("pro == %s", pro)
	return pro + " to User json"
}

func (this *XXAgent) XXConvertAudit(pro string) string {
	fmt.Printf("pro222 == %s", pro)
	return pro + " to Audit json"
}

func main() {

	var agent XXAgent

	//用于通过字符串调用相应的函数
	cMap := make(ControlMapType, 0)

	//创建反射变量 ,需要传入变量的地址，否则后续只能反射 静态方法
	ref := reflect.ValueOf(&agent)
	vft := ref.Type()

	//读取 agent 绑定的方法数量
	num := vft.NumMethod()
	fmt.Println("num == ", num)

	for i := 0; i < num; i++ {
		//fmt.Println(ref.Method(i))      //具体函数的的地址
		//fmt.Println(vft.Method(i).Type) // Method结构体，其中成员有Name，Pkgpath，Type，Func，Index

		name := vft.Method(i).Name
		fmt.Println("name == ", name)
		cMap[name] = ref.Method(i)
	}

	info := "i am proto"
	//传入参数列表
	parms := []reflect.Value{reflect.ValueOf(info)}
	res := cMap["XXConvertUser"].Call(parms)
	fmt.Printf("\nres == %T\n", res)
}
