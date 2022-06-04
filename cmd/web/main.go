package main

import (
	"flag"
	"log"
	"net/http"
	"path/filepath"
)

func main() {

	// create a new flag cmd
	addr := flag.String("addr", ":4000", "Net address HTTP")

	// call func for extract flag from cmd
	flag.Parse()

	// initialisation new router
	mux := http.NewServeMux()

	// registaration handlers
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)


	// handle http-requests to static files from "/static"
	// fileServer := http.FileServer(neuteredFileSystem{http.Dir("C:/Users/Lesik/go/src/snippetbox/ui/static")})
	fileServer := http.FileServer(http.Dir("C:/Users/Lesik/go/src/snippetbox/ui/static"))

	// use for registration handle all requests, begining with "/static/"
	// mux.Handle("/static", http.NotFoundHandler())
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// uses ListenAndServe to run a new server
	// we hang over two parameters : TCP adress for listening, and created router
	log.Printf("Web server start on %s", *addr)

	err := http.ListenAndServe(*addr, mux)

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
