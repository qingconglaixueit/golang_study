package main

import (
	"fmt"
	"sync"
)

//sync.Once 只会执行一次，源码是用互斥锁和atomic 加载变量 和 存储变量来做的
func main() {
	var one sync.Once

	one.Do(pr)

	one.Do(pr)

	one.Do(pr)
}

func pr() {
	fmt.Print("i am pr")
}
