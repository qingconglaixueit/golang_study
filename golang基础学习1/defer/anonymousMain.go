package main

import "fmt"

func main() {
	// a := 10
	// b := 20

	// defer func() {
	// 	fmt.Printf("inner  a == %d, b == %d\n", a, b)
	// }()

	// a = 1
	// b = 2
	// fmt.Printf("externer  a == %d, b == %d\n", a, b)

// externer  a == 1, b == 2
// inner  a == 1, b == 2

	a := 10
	b := 20

	defer func(a, b int) {
		fmt.Printf("inner  a == %d, b == %d\n", a, b)
	}(a, b) //此处参数 a=10 b=20的参与已经传入到函数中了，只是匿名函数最后执行而已

	a = 1
	b = 2
	fmt.Printf("externer  a == %d, b == %d\n", a, b)

// externer  a == 1, b == 2
// inner  a == 10, b == 20
}
