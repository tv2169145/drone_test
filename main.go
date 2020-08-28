package main

import (
	"io"
	"net/http"
)

const (
	name = "jimmy"
)

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {

	io.WriteString(w, GetHelloWorld())
}

func GetHelloWorld() string {
	return "hello " + name + " 123456"
}


