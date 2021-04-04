package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print(runtime.GOMAXPROCS(0))

	runtime.GOMAXPROCS(10)


	fmt.Print(runtime.GOMAXPROCS(0))

}
