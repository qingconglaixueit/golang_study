package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type test struct {
	Name  string
	Hobby string
}

func myHandle(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	par := req.URL.Query()
	fmt.Println("par :", par)
	//回写数据
	resp := &test{
		Name:  par.Get("name"),
		Hobby: par.Get("hobby"),
	}
	respByte, _ := json.Marshal(resp)
	fmt.Fprintln(w, string(respByte))

}

//server端
func main() {

	http.HandleFunc("/", myHandle)

	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		fmt.Printf("ListenAndServe err : %v", err)
		return
	}

}
