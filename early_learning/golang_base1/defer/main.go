package main

import "fmt"

func test(x int) int {
	res := 100 / x
	return res
}

func main() {
	// fmt.Println("hello")

	// defer fmt.Println("over!!")

	// fmt.Println("world")

	// //多个defer的执行顺序，类似于栈的方式，先进后出，哪怕中间出现程序错误，defer修饰的语句仍然会执行
	// for i := 0; i < 10; i++ {
	// 	defer fmt.Println(i)
	// }

	defer fmt.Println("aaaa")
	defer fmt.Println("bbb")
	defer test(0)
	defer fmt.Println("ccc")

}
