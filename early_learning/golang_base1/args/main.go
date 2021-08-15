package main

import "fmt"

func getname(args ...string) {
	length := len(args)
	if length > 0 {
		for i, data := range args {
			fmt.Printf("%d -- %s\n", i, data)
		}
	}
}

func test(args ...string) {
	getname(args[1:]...)
}

func main() {
	test("qqq", "yyy", "bbb")
}
