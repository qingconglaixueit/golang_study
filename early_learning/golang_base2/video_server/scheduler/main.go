package main

import (
	"github.com/julienschmidt/httprouter"
	"video_server/scheduler/taskrunner"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.GET("/video-delete-record/:vid-id", vidDelRecHandler)
	return router
}

func main() {
	go taskrunner.Start()
	select {}
	//r := RegisterHandlers()
	//http.ListenAndServe(":10001", r)
}
