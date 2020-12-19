package main

import "fmt"

func Add(a, b int) int {
	return a + b
}

type Universal func(int, int) int

func main() {

	var uni Universal
	uni = Add
	res := uni(1, 2)
	fmt.Println(res)

}
