package main

import (
	"log"
	"net/http"
)

// handler
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from snippetbox"))
}

func main() {

	// to initialize a new servemux
	mux := http.NewServeMux()
	// register the home function as the handler for the "/" URL pattern
	mux.HandleFunc("/", home)

	log.Println("starting server on: 4000")
	// http.ListenAndServe() to start
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}
