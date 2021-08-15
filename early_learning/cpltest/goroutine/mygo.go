package main

import "fmt"

func Hi() {
	fmt.Println("this is Hi Goroutine!")
}
func main() {
	go Hi()
	fmt.Println("main goroutine!")
}
