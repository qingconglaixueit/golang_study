package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		defer func() {
			fmt.Println("run defer")
			if err := recover(); err != nil {
				fmt.Println("catch panic , is recovering")
			}
		}()

		fmt.Println("run child routine")
		var ptr *int
		*ptr = 0x123456
	}()
	time.Sleep(time.Second * time.Duration(1))
}
