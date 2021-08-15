package main

import (
	"encoding/json"
	"fmt"
)


func main(){
	// 是一个空对象
	var mys1 []int
	// 是一个对象，对象里面是一个切片，这个切片没有元素
	var mys2 = []int{}

	json1, _ := json.Marshal(mys1)
	json2, _ := json.Marshal(mys2)

	fmt.Println(string(json1))
	fmt.Println(string(json2))
}


