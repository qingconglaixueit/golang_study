package main

import "fmt"

//func main(){
//
//	mys := make([]int,3,5)
//	fmt.Println("len == ", len(mys))
//	fmt.Println("cap == ", cap(mys))
//
//	mys[0] = 1
//	mys[1] = 1
//	mys[2] = 1
//
//	mys = append(mys,2)
//
//	fmt.Println("len == ", len(mys))
//	fmt.Println("cap == ", cap(mys))
//	for _,v :=range mys{
//		fmt.Printf("%v",v)
//	}
//}

func main(){

	arr := [8]int{}

	mySlice := arr[4:6]

	fmt.Println("len == ", len(mySlice))
	fmt.Println("cap == ", cap(mySlice))

	// 此处的遍历 长度是 len 的长度
	for _,v :=range mySlice{
		fmt.Printf("%v",v)
	}
}