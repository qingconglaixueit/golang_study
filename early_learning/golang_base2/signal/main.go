package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

//将信号放到通道里面，调用系统的信号通知函数进行提前注册
func main(){

	sig := make(chan os.Signal,1)
	signal.Notify(sig,syscall.SIGINT)
	fmt.Print("wait ctrl + C")
	fmt.Printf("signal is %v",<-sig)

}
