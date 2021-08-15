package main

import "fmt"

func OnlyWriteData(out chan<- int) {
	// 单向 通道 ， 只写 不能读
	for i := 0; i < 10; i++ {
		out <- i
	}
	close(out)
}

func CalData(out chan<- int, in <-chan int) {
	// out 单向 通道 ， 只写 不能读
	// int 单向 通道 ， 只读 不能写

	// 遍历 读取in 通道，若 in通道 数据读取完毕，则阻塞，若in 通道关闭，则退出循环
	for i := range in {
		out <- i + i
	}
	close(out)
}
func myPrinter(in <-chan int) {
	// 遍历 读取in 通道，若 in通道 数据读取完毕，则阻塞，若in 通道关闭，则退出循环
	for i := range in {
		fmt.Println(i)
	}
}

func main() {
	// 创建2 个无缓冲的通道
	ch1 := make(chan int)
	ch2 := make(chan int)


	go OnlyWriteData(ch1)
	go CalData(ch2, ch1)


	myPrinter(ch2)
}