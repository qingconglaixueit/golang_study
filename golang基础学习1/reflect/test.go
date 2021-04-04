package main

import "fmt"

func modifyName(name *string) {
	*name = "world"
}

func main() {
	name := "hello"

	fmt.Println(name)

	modifyName(&name)

	fmt.Println(name)

}
