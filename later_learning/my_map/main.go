package main

import (
	"encoding/json"
	"log"
)

func main(){

	log.SetFlags(log.Ldate | log.Ltime| log.Llongfile)
	m  := make(map[int]int)
	m[0] = 1
	log.Println(m)

	jsonB, _ :=json.Marshal(m)

	unM := make(map[int]int)
	_ = json.Unmarshal(jsonB,&unM)

	log.Println(unM)
}
