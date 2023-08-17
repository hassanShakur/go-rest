package main

import (
	"io"
	"log"
	"net/http"
)

func MyServer (w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello world\n")
}

func main () {
	http.HandleFunc("/hello", MyServer)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
