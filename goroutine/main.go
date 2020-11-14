package main

import (
	"fmt"
	"time"
)

func say(s string) {
	var i int
	for i = 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

var num int = 0

//goroutine 是线程不安全的
func countNum() {
	var i int
	for i = 0; i < 10; i++ {
		time.Sleep(5 * time.Millisecond)
		num++
	}
}
func main() {
	//go say("hello")
	//say("world")

	go countNum()
	countNum()
	fmt.Println(num)
}
