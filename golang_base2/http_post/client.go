package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

//post client

func main() {

	reqUrl := "http://127.0.0.1:8888"
	contentType := "application/json"
	data := `{"name":"qing","age":18}`

	resp, err := http.Post(reqUrl, contentType, strings.NewReader(data))
	if err != nil {
		fmt.Printf("Post err %v", err)
		return
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("ReadAll err %v", err)
		return
	}

	fmt.Println(string(b))

}
