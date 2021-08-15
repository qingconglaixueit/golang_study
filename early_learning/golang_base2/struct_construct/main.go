package main

import "fmt"

type phone struct{
	band string
	price int
}
//自定义构造函数，工厂模式
func NewPhone(band string,price int)*phone{
	return &phone{band,price}
}

//&{xiaomi 2000}
//&{huawei 3000}
func main(){
	fmt.Println(NewPhone("xiaomi",2000))
	fmt.Println(NewPhone("huawei",3000))

}