package main

import "fmt"

//不带缓冲的 通道
func getSum(s []int, c chan int) {
	sum := 0
	for _, value := range s {
		sum += value
	}
	c <- sum
}

func getSum2(c chan int, n int) {
	x, y := 0, 1
	var i int
	for i = 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c) //关闭通道
}

func main() {
	//不带缓冲的 通道
	// s := []int{3, 5, -2, 3, 4, 7, 1, 1, 1}
	// c := make(chan int)
	// go getSum(s[:3], c)
	// go getSum(s[3:6], c)
	// go getSum(s[6:], c)
	// x, y, z := <-c, <-c, <-c
	// fmt.Println(x, y, z, x+y+z)

	//带缓冲的通道
	c := make(chan int, 10)
	go getSum2(c, cap(c))

	for value := range c {
		fmt.Println(value)
	}

}
