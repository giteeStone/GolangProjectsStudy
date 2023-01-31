package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	myhandler := Myhandler{}
	server := http.Server{
		Addr:        ":8888",
		Handler:     &myhandler,
		ReadTimeout: 2 * time.Second,
	}
	server.ListenAndServe()
}

type Myhandler struct{}

func (h *Myhandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "server")
}
