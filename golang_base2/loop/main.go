package main

import (
	"fmt"
	"time"
)
//如果有个别协程死循环了会导致其它协程饥饿得到不运行么？  不会
func main() {

	n := 3
	for i := 0; i < n; i++ {
		go func() {
			fmt.Printf("run %d child goroutine", i)

			for {

			}
		}()
	}

	for {
		time.Sleep(time.Second)
		fmt.Print("run main goroutine")
	}

}
