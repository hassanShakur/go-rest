package main

import (
	"fmt"
	"net/http"
)

func MultiHandler() {
	newMux := http.NewServeMux()

	newMux.HandleFunc("/abc", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "The abc route")
	})

	newMux.HandleFunc("/123", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "The 123 route")
	})

	http.ListenAndServe(":8000", newMux)
}