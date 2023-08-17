# Go Rest-ful

## Simple Server & Router

```go

// A struct which can be a multiplexer
type MyServeMux struct {}

// This is the function handler to be overridden
// Overrides the http.HandleFunc which is executed in simple routing
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
```

## Different Handlers

With different handlers for different routes:

```go
newMux := http.NewServeMux()

	newMux.HandleFunc("/abc", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "The abc route")
	})

	newMux.HandleFunc("/123", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "The 123 route")
	})
```
