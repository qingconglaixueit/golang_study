package main

import (
	"flag"
	"fmt"
)

func main() {
	//    fmt.Println(os.Args)
	ok := flag.Bool("ok", false, "is ok") // 不设置ok 则为false
	ok2 := flag.Bool("ok2", false, "is okss") // 不设置ok2 则为false
	id := flag.Int("id", 0, "id")
	port := flag.String("port", ":8080", "http listen port")
	var name string
	flag.StringVar(&name, "name", "Jack", "name")

	flag.Parse()
	//flag.Usage()
	others := flag.Args()

	fmt.Println("ok:", *ok)
	fmt.Println("ok2:", *ok2)
	fmt.Println("id:", *id)
	fmt.Println("port:", *port)
	fmt.Println("name:", name)
	fmt.Println("other:", others)
}
