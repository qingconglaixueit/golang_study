package main

import (
	"fmt"
	"github.com/robfig/cron"
	"log"
)

func main() {
	i := 0
	c := cron.New()
	spec := "*/5 * * * * ?"
	err:=c.AddFunc(spec, func() {
		i++
		log.Println("cron running:", i)
	})
	fmt.Println(err)
	c.Start()

	select{}
}