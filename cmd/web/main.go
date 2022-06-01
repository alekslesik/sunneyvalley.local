package main

import (
	"log"
	"net/http"
	"path/filepath"
)

func main() {

	// initialisation new router
	mux := http.NewServeMux()

	// registaration handlers
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)


	// handle http-requests to static files from "/static"
	fileServer := http.FileServer(neuteredFileSystem{http.Dir("C:/Users/Lesik/go/src/snippetbox/ui/static")})

	// use for registration handle all requests, begining with "/static/"
	mux.Handle("/static", http.NotFoundHandler())
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// uses ListenAndServe to run a new server
	// we hang over two parameters : TCP adress for listening, and created router
	log.Println("Web server start on 127.0.0.1:4000")

	err := http.ListenAndServe(":4000", mux)

	log.Fatal(err)
}

type neuteredFileSystem struct {
	fs http.FileSystem
}

func (nfs neuteredFileSystem) Open(path string) (http.File, error) {
	f, err := nfs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if s.IsDir() {
		index := filepath.Join(path, "index.html")
		if _, err := nfs.fs.Open(index); err != nil {
			closeEr := f.Close()
			if closeEr != nil {
				return nil, closeEr
			}

			return nil, err
		}
	}

	return f, nil
}
