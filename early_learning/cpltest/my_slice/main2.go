package main

import "fmt"

func main(){

	mySlice := make([]int,4,4)
	mySlice[0] = 3
	mySlice[1] = 6
	mySlice[2] = 7
	mySlice[3] = 8

	fmt.Printf("ptr == %p\n", &mySlice)
	fmt.Println("len == ", len(mySlice))
	fmt.Println("cap == ", cap(mySlice))

	// 此处的遍历 长度是 len 的长度
	for _,v :=range mySlice{
		fmt.Printf("%v ",v)
	}

	fmt.Println("")


	mySlice = append(mySlice,5)

	fmt.Printf("new_ptr == %p\n", &mySlice)
	fmt.Println("new_len == ", len(mySlice))
	fmt.Println("new_cap == ", cap(mySlice))

	// 此处的遍历 长度是 len 的长度
	for _,v :=range mySlice{
		fmt.Printf("%v ",v)
	}
}

