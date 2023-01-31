package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	myhandler := Myhandler{}
	mux.Handle("/demo", &myhandler)
	server := http.Server{
		Addr:        ":8888",
		Handler:     mux,
		ReadTimeout: 2 * time.Second,
	}
	server.ListenAndServe()
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "helloHandler")
}

type Myhandler struct{}

func (h *Myhandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "server myHandler")
}
