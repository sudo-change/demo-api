package main

import (
	"log"
	"net/http"
)

// handler
func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from snippetbox"))
}

func main() {

	// to initialize a new servemux
	mux := http.NewServeMux()
	// register the home function as the handler for the "/" URL pattern
	mux.HandleFunc("/", home)

	log.Println("starting server on: 4000")
	// http.ListenAndServe() to  a new web server
	// if it returns an error, we use log.Fatal() function to log the error message and exit
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}
