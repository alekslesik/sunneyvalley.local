package main

import (
	"log"
	"net/http"
)

func main() {

	// initialisation new router
	mux := http.NewServeMux()

	// registaration handlers
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	// uses ListenAndServe to run a new server
	// we hang over two parameters : TCP adress for listening, and created router
	log.Println("Web server start on 127.0.0.1:4000")

	err := http.ListenAndServe(":4000", mux)

	log.Fatal(err)

}