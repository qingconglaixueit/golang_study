package main

import (
	"fmt"
	"sync"
)

// 全局变量
var num int64
var wg sync.WaitGroup
var lock sync.Mutex

func add() {
	for i := 0; i < 10000000; i++ {
		// 访问资源前  加锁
		lock.Lock()
		num = num + 1
		// 访问资源后  解锁
		lock.Unlock()
	}
	// 协程退出， 记录 -1
	wg.Done()
}
func main() {
	// 启动2个协程，记录 2
	wg.Add(2)

	go add()
	go add()

	// 等待子协程退出
	wg.Wait()
	fmt.Println(num)
}
