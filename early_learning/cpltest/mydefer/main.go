package main

import "fmt"

func test3() (res int) {
	defer func() {
		res++
	}()

	return 1
}
func main() {

	fmt.Println(test3())

	return
}
