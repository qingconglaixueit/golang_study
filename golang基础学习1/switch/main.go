package main

import (
	"fmt"
)

func main() {

	num := 2
	switch num {
	case 1:
		fmt.Println("111")
	case 2:
		fmt.Println("222")
		fallthrough
	case 3:
		fmt.Println("333")
	case 4:
		fmt.Println("4444")
	}

	score := 90
	switch { //这里可以不用写 条件
	case score > 90: //case 后面可以写条件
		fmt.Println("extent")
	case score > 80:
		fmt.Println("good")
	case score > 70:
		fmt.Println("not bad")
	}

}
