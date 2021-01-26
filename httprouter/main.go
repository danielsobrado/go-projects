package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func index(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "Index")
}

func hello(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", p.ByName("name"))
}

func world(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "World!")
}

func main() {

	mux := httprouter.New()
	mux.GET("/hello/:name", hello)
	mux.GET("/index", index)
	mux.GET("/world", world)

	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}

	server.ListenAndServe()
}
