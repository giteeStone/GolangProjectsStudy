package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", hello)
	myhandler := Myhandler{}
	http.Handle("/demo", &myhandler)
	http.ListenAndServe(":8888", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello wolrd", r.URL.Path)
}

type Myhandler struct{}

func (h *Myhandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "demo!!!!!!")
}
