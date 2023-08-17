package main
import (
 "fmt"
 "math/rand"
 "net/http"
)


type MyServeMux struct {}


func (p *MyServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		giveRandom(w, r)
		return
	}

	http.NotFound(w, r)
}


func giveRandom(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Your random number is: %f", rand.Float64())
}


func MyMain() {
	// Any struct that has serveHTTP function can be a multiplexer
	mux := &MyServeMux{}
	http.ListenAndServe(":8000", mux)
}
