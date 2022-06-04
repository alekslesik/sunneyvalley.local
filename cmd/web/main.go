package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {

	// create a new flag cmd
	addr := flag.String("addr", ":4000", "Net address HTTP")

	// call func for extract flag from cmd
	flag.Parse()

	f, err := os.OpenFile("info.log", os.O_RDWR | os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// create logger for writing info messages
	infoLog := log.New(f, "INFO\t", log.Ldate | log.Ltime)

	// create logger for writing error messages
	errorLog := log.New(f, "ERROR\t", log.Ldate | log.Ltime | log.Lshortfile)

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

	// initialise new struct, set fields Addr and Handler
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	// apply created loggers
	infoLog.Printf("Web server start on %s", *addr)

	// uses ListenAndServe to run a new server
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
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
