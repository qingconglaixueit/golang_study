package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	list := os.Args
	fmt.Println(len(list))
	fmt.Println(list)
	cmd := exec.Command("bash", "-c", `ls -al`)
	fmt.Println(cmd.Start())
}
