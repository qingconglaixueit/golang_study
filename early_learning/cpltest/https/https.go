package main

import (
	"log"
	"net/http"

	"github.com/unrolled/secure"
)

var myHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>hello xiaomotong!! HTTPS </h1>"))
})

func main() {
	secureMiddleware := secure.New(secure.Options{
		SSLRedirect: true,
		// 这在生产中是可选的。默认行为是将请求重定向到 HTTPS 协议
		SSLHost:     "localhost:4433",
	})

	app := secureMiddleware.Handler(myHandler)

	// HTTP
	go func() {
		log.Fatal(http.ListenAndServe(":8888", app))
	}()

	log.Fatal(http.ListenAndServeTLS(":4433", "cert.pem", "key.pem", app))
}
