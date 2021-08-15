package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			// 循环向无缓冲的通道中写入数据， 只有当上一个数据被读走之后，下一个数据才能往通道中放
			c <- i
		}
		// 关闭通道
		close(c)
	}()
	for {
		// 读取通道中的数据，若通道中无数据，则阻塞，若读到 nil， 则通道关闭，退出循环
		if data, ok := <-c; ok {
			fmt.Println(data)
		} else {
			break
		}
	}
	fmt.Println("channel over")
}
