package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func handPost(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	if req.Method == http.MethodPost {
		b, err := ioutil.ReadAll(req.Body)
		if err != nil {
			fmt.Printf("ReadAll err %v", err)
			return
		}

		fmt.Println(string(b))

		resp := `{"status":"200 OK"}`

		w.Write([]byte(resp))

		fmt.Println("reponse post func")
	} else {
		fmt.Println("can't handle ", req.Method)
		w.Write([]byte(http.StatusText(http.StatusBadRequest)))
	}
}

//post server

func main() {

	http.HandleFunc("/", handPost)

	err := http.ListenAndServe("0.0.0.0:8888", nil)
	if err != nil {
		fmt.Printf("ListenAndServe err %v", err)
		return
	}
}
