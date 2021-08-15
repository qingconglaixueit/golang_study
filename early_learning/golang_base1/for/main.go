package main

import "fmt"

func main() {

	str := []string{"北京", "天津", "山东"}

	for i := range str {
		fmt.Printf("%d -- %s\n", i, str[i])
	}

}
