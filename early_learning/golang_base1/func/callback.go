package main

import "fmt"

func Add(a, b int) int {
	return a + b
}

type Universal func(int, int) int

func cal(uni Universal, a int, b int) int {
	return uni(a, b)
}
func main() {
	uni := Add
	res := cal(uni, 1, 2)
	fmt.Println(res)
}
