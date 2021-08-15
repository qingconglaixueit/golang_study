package main

import (
	"fmt"
	"runtime"
	"time"
)

const N = 1000000
//模拟开百万协程，然而 开40w的协程 就出现问题了
func main() {
	fmt.Print("run in main goroutine")
	var i = 1
	for {
		go func() {
			for {
				time.Sleep(time.Second)
			}
		}()
		i++
		if i%10000 == 0 {
			fmt.Println("%d  start goroutine", i)
		}

		if i == N {
			break
		}
	}
	fmt.Println("NumGoroutine:", runtime.NumGoroutine())
	time.Sleep(time.Second * 15)

}
