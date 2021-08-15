package main

import (
	"fmt"
	"time"
)

//定义函数类型
type Fn func()

//定时器中的成员
type MyTicker struct {
	MyTick *time.Ticker
	Runner Fn
}

func NewMyTick(interval int, f Fn) *MyTicker {
	return &MyTicker{
		MyTick: time.NewTicker(time.Duration(interval) * time.Second),
		Runner: f,
	}
}

//启动定时器需要执行的任务
func (t *MyTicker) Start() {
	for {
		select {
		case <-t.MyTick.C:
			t.Runner()
		}
	}
}

func testChannelTimeout(conn chan int) bool {
	// 设置 1 秒的定时器，若在到了1 s ,则进行打印，说明已经超时
	timer := time.NewTimer(1 * time.Second)

	select {
	case <-conn:
		if timer.Stop() {
			fmt.Println("timer.Stop()")
		}
		return true
	case <-timer.C: // timer 通道超时
		fmt.Println("timer Channel timeout!")
		return false
	}
	//ch := make(chan int, 1)
	//ch <- 1
	//go testChannelTimeout(ch)
}
func testPrint(){
	fmt.Println(" 滴答 1 次")
}
func main() {

	//t := NewMyTick( 1 ,testPrint)
	//t.Start()

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	// 若通道为空，则阻塞
	// 若通道有数据，则读取
	// 若通道关闭，则退出
	for range ticker.C {
		fmt.Println("ticker ticker ticker ...")
	}
}
