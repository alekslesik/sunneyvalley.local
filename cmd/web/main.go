package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"golangs.org/snippetbox/pkg/config"
)

func main() {
	// log file
	f, err := os.OpenFile("info.log", os.O_RDWR | os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// set config.Application
	app := &config.Application{
		ErrorLog: log.New(f, "ERROR\t", log.Ldate | log.Ltime | log.Lshortfile),
		InfoLog: log.New(f, "INFO\t", log.Ldate | log.Ltime),
	}

	// initialisation new router
	mux := http.NewServeMux()
	// registaration handlers
	mux.HandleFunc("/", Home(app))
	mux.HandleFunc("/snippet", ShowSnippet(app))
	mux.HandleFunc("/snippet/create", CreateSnippet(app))

	// handle http-requests to static files from "/static"
	// fileServer := http.FileServer(neuteredFileSystem{http.Dir("C:/Users/Lesik/go/src/snippetbox/ui/static")})
	fileServer := http.FileServer(http.Dir("C:/Users/Lesik/go/src/snippetbox/ui/static"))

	// use for registration handle all requests, begining with "/static/"
	// mux.Handle("/static", http.NotFoundHandler())
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// create a new flag cmd
	addr := flag.String("addr", ":4000", "Net address HTTP")
	// call func for extract flag from cmd
	flag.Parse()

	// initialise new struct, set fields Addr and Handler
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: app.ErrorLog,
		Handler:  mux,
	}

	// apply created loggers
	app.InfoLog.Printf("Web server start on %s", *addr)

	// uses ListenAndServe to run a new server
	err = srv.ListenAndServe()
	app.ErrorLog.Fatal(err)
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
