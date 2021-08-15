package main

import "log"

//可变参数

func add(a int, arg ...int) int {

	var tmp int
	for _, v := range arg {
		tmp += v
	}
	return tmp + a

}

func main() {

	log.Println("add resault == ", add(1, 2, 3, 4, 5, 6, 7))

}
