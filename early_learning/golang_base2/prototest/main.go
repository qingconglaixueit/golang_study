package main

import (
	"acagent/protoc/all"
	"acagent/protoc/myurl"
	"fmt"
	"github.com/golang/protobuf/proto"
)

func main() {

	urlLib := &myurl.UrlType{
		AccountId: "111111",
		Name:      "xiaozhu",
		Des:       "miaoshu",
		Keywords:  "guanjianzi",
	}

	res, err := proto.Marshal(urlLib)
	if err != nil{
		fmt.Println("Marshal failed!!")
		return
	}

	allData := &all.AllType{}

	err = proto.Unmarshal(res,allData)
	if err != nil{
		fmt.Println("Unmarshal failed")
		return
	}

	fmt.Println(allData)





}
