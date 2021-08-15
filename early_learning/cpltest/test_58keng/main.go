package main

import "runtime"

type data struct {
	Name string
}

func main() {
	done := false

	go func() {
		done = true
	}()

	for !done {
		runtime.Gosched()
	}

	println("done !")
}
