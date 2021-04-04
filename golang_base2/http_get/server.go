package main

import (
	"fmt"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("hello world"))
}

func sayQing(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("hello qingbing"))
}

//server端
func main() {

	//sm := http.NewServeMux()
	//
	//sm.HandleFunc("/", sayHello)
	//sm.HandleFunc("/qing", sayQing)
	//
	//err := http.ListenAndServe("0.0.0.0:8888", sm)

	http.HandleFunc("/", sayHello)
	http.HandleFunc("/qing", sayQing)

	err := http.ListenAndServe("0.0.0.0:8888", nil)
	if err != nil {
		fmt.Printf("ListenAndServe err: %v", err)
		return
	}
}
