package main

import (
	"github.com/wonderivan/logger"
	"net/http"
	"time"

	// "html/template"
	"github.com/julienschmidt/httprouter"
)

type middleWareHandler struct {
	r *httprouter.Router
}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t1 := time.Now().UnixNano() / 1e6
	logger.Info("ServeHTTP =========")
	m.r.ServeHTTP(w, r)
	logger.Info("ServeHTTP =========")
	t2 := time.Now().UnixNano() / 1e6
	logger.Info("Method:%s, url:%s need time:%dms\n\n", r.Method, r.URL.Path, t2-t1)
}

func NewMiddleWareHandler(r *httprouter.Router) http.Handler {
	m := middleWareHandler{}
	m.r = r
	return m
}
func RegisterHandler() *httprouter.Router {
	router := httprouter.New()
	router.GET("/", homeHandler)
	router.POST("/", homeHandler)
	router.GET("/userhome", userHomeHandler)
	router.POST("/userhome", userHomeHandler)
	router.POST("/api", apiHandler)              // api手动转发
	router.POST("/upload/:vid-id", proxyHandler) // 上传文件的时候
	router.ServeFiles("/statics/*filepath", http.Dir("./templates"))
	router.ServeFiles("/scripts/*filepath", http.Dir("./templates/scripts"))
	return router
}

func main() {

	logger.SetLogger("./log.json")

	r := RegisterHandler()
	mh := NewMiddleWareHandler(r)
	http.ListenAndServe(":10003", mh)
}
