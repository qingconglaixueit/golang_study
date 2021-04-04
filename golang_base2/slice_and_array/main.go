package main

import "fmt"

//[1 2 3] [1 2 3]
//[1 2 3] [999 2 3]
func array_test() {
	var x [3]int = [3]int{1, 2, 3}
	var y [3]int = x
	fmt.Println(x, y)
	y[0] = 999
	fmt.Println(x, y)
}

//slice，map，chan 是引用类型
//[1 2 3] [1 2 3]
//[999 2 3] [999 2 3]
func slice_test() {
	var x []int = []int{1, 2, 3}
	var y []int = x
	fmt.Println(x, y)
	y[0] = 999
	fmt.Println(x, y)
}

func main() {
	array_test()
	fmt.Sprint("===================")
	slice_test()
}
