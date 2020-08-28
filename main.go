package main

import (
	"github.com/tv2169145/drone_test/utils"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {

	io.WriteString(w, utils.GetHelloWorld())
}


