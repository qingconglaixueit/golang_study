package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var num int64
var l sync.Mutex
var wg sync.WaitGroup

// 普通版加函数
func add() {
	num = num + 1
	wg.Done()
}

// 互斥锁版加函数
func mutexAdd() {
	l.Lock()
	num = num + 1
	l.Unlock()
	wg.Done()
}

// 原子操作版加函数
func atomicAdd() {
	atomic.AddInt64(&num, 1)
	wg.Done()
}

func main() {
	// 目的是 记录程序消耗时间
	start := time.Now()
	for i := 0; i < 20000; i++ {

		wg.Add(1)

		// go add()       // 无锁的  add函数 不是并发安全的
		 go mutexAdd()  // 互斥锁的 add函数 是并发安全的，因为拿不到互斥锁会阻塞，所以加锁性能开销大

		//go atomicAdd()    // 原子操作的 add函数 是并发安全，性能优于加锁的
	}

	// 等待子协程 退出
	wg.Wait()

	end := time.Now()
	fmt.Println(num)
	// 打印程序消耗时间
	fmt.Println(end.Sub(start))
}