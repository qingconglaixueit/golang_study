package main

import (
	"fmt"
	"time"
)

func main() {

	gj := make(chan string, 1)

	ch := make(chan bool, 1)
	ch <- false
	fmt.Println("hello world")

	go func() {
		var flag bool
		var times = 0
		for {
			time.Sleep(time.Second * time.Duration(1))
			times++

			if times > 5{
				gj <- "告警：实时下发数据 出错！！！！！，超过5次了"
			}

			fmt.Println("line 18 --- 实时收到 pop下发数据")
			select {
			case flag = <-ch:
				if flag == true {
					fmt.Println("line 22 --- 服务正在处理数据，很忙...")
				} else {
					fmt.Println("现在服务空闲，可以处理数据...")
					ch <- true
					for i := 0; i < 3; i++ {
						time.Sleep(time.Second * time.Duration(1))
						fmt.Println("line 28 ---数据正在处理  ", (float32(i)+1.0)/3.0)
					}
					fmt.Println("line 30 --- 数据处理完毕  ")
					ch <- false
				}
			}
		}

	}()

	go func() {
		var ff bool
		var times = 0
		for {
			time.Sleep(time.Second * time.Duration(2))
			times++
			if times > 3{
				gj <- "告警：拉取数据 出错！！！！！，超过5次了"
			}
			fmt.Println("line 41 --- 准备拉取数据...")
			select {
			case ff = <-ch:
				if ff == true {
					fmt.Println("line 45 --- 还不能拉取数据，很忙...")
				} else {
					ch <- true
					fmt.Println("line 48 --- 现在空闲，正在拉取数据...")
					for i := 0; i < 3; i++ {
						time.Sleep(time.Second * time.Duration(1))
						fmt.Println("line 51 --- 正在拉取数据  ", (float32(i)+1.0)/3.0)
					}
					fmt.Println("line 53 --- 拉取完毕")
					ch <- false
				}
			}
		}
	}()

	var info string
	for{
		select {
		case info = <-gj:
			fmt.Printf("收到告警信息 ---- %s , 需要人工介入\n",info)

		}
	}

}
