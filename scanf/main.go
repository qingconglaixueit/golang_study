package main

import "fmt"

func main() {
	var tmp int
	fmt.Scanf("%d", &tmp)
	fmt.Println("tmp == ", tmp)

	var tmp2 int
	fmt.Scan(&tmp2)
	fmt.Printf("type == %T", tmp2)
}
