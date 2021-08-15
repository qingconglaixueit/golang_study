package main

import (
	"context"
	"fmt"
	"time"
)

var key1 = "hello"
var key2 = "hello2"
func main() {

	//ctx1, cancel1 := context.WithTimeout(context.Background(), 2)
	ctx, cancle := context.WithCancel(context.Background())
	ctx1 := context.WithValue(ctx, key1, "xaohuo")
	ctx2 := context.WithValue(ctx1, key2, "dahuo")


	go func(ctx2 context.Context) {
		for {
			select {
			case <-ctx2.Done():
				fmt.Println("ctx1 done", ctx2.Err(), ctx2.Value(key2))
				return
			default:
				time.Sleep(time.Second * 2)
				fmt.Println("监控中")
				time.Sleep(time.Second * 2)
			}
		}

	}(ctx2)

	time.Sleep(time.Second * 3)
	fmt.Println("现在执行cancle1")
	cancle()
	time.Sleep(time.Second * 10)


}
