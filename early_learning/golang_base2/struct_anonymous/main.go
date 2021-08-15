package main

import "fmt"

//匿名变量，赋值或者初始化的时候，字段名字和结构体名字一致才行

type Person struct {
	Name string
	Age  int
}
type Student struct {
	Hobby string
	Person

}

func main() {

	stu := &Student{
		Hobby: "play basketball",
	}
	stu.Person.Age = 16
	stu.Person.Name = "xiaozhu"

	stu.Name = "hhh"
	stu.Age = 18


	fmt.Print(stu)

}
