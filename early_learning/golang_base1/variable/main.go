package main

import (
	"fmt"
	"hhh/test"
)

var a int = 10

func init() {
	fmt.Println("this is main init")
}

func main() {

	a := "xiaozhu"
	fmt.Println(a)
	{
		a := 1.4
		fmt.Println(a)
	}
	test.Pr()
}
