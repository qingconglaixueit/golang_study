package main

import (
	"fmt"
	"time"
)

//制造panic
func main() {
	fmt.Print("enter main\n")
	go func() {
		fmt.Print("enter child 1\n")
		go func() {
			fmt.Sprint("enter grand child\n")
			go func() {
				defer func() {
					if err := recover(); err != nil {
						fmt.Println(err)
					}
				}()
				fmt.Println("enter grand grand child")
				var ptr *int
				*ptr = 0x123456
				time.Sleep(time.Second * time.Duration(1))
				fmt.Println("quit grand grand  grand child")
			}()
			time.Sleep(time.Second * time.Duration(1))
			fmt.Println("quit grand child")
		}()
		time.Sleep(time.Second * time.Duration(1))
		fmt.Println("quit child")
	}()
	time.Sleep(time.Second * time.Duration(1))
	fmt.Println("quit main")
}
