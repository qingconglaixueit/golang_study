package main

import "fmt"

func main() {
	type long int64

	var a long
	a = 2

	fmt.Printf("type == %T", a)

}
