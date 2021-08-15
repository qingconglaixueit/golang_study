package main

import "fmt"

func cal() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func main() {

	//1、闭包的特点1，可以捕获变量，且是以引用的方式捕获的
	a := 1
	name := "go"
	func() {
		a = 2
		name = "hello wrold"
		fmt.Printf("内部  ： a == %d, name == %s\n", a, name)
	}()
	fmt.Printf("外部  ： a == %d, name == %s\n", a, name)

	//2、必要的特点2 ，闭包里面使用的变量，只要闭包还在使用，则一直有效，哪怕超出该变量的作用域
	f := cal()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}
