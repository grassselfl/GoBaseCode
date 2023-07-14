package main

import (
	"fmt"
	"net/http"
)

func f1(w http.ResponseWriter, r *http.Request) {
	write, err := w.Write([]byte(`<h1 style="color:red">Hello World<h1>`))
	if err != nil {
		return
	}
	fmt.Println(write)

	fmt.Printf("%v", r)
	fmt.Println(r.Method)
	fmt.Println(r.Host)
	fmt.Println(r.RequestURI)
}

func main() {
	http.HandleFunc("/api/hello", f1)
	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		return
	}
}
