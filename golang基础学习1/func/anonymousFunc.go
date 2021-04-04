package main

import "fmt"

func main() {

	name := "qqq"
	age := 24
	//1、匿名函数
	f1 := func() {
		fmt.Printf("name == %s\n", name)
		fmt.Printf("age == %d\n", age)
	}

	f1()

	//2、定义匿名函数的时候直接调用
	func(hobby string) {
		fmt.Println(name)
		fmt.Println(hobby)
	}("basketball")

	//3、定义匿名函数有参数有返回值
	f3 := func(a, b int) int {
		return a + b
	}

	fmt.Println(f3(1, 3))

	res := func(a, b int) int {
		return a + b
	}(3, 8)
	fmt.Println(res)

}
