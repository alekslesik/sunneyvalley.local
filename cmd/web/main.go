package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {

	// create a new flag cmd
	addr := flag.String("addr", ":4000", "Net address HTTP")
	// call func for extract flag from cmd

	flag.Parse()
	// log file
	f, err := os.OpenFile("info.log", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// set application
	app := application{
		errorLog: log.New(f, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
		infoLog:  log.New(f, "INFO\t", log.Ldate|log.Ltime),
	}

	// initialise new struct, set fields Addr and Handler
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: app.errorLog,
		Handler:  app.routes(),
	}

	// apply created loggers
	app.infoLog.Printf("Web server start on %s", *addr)

	// uses ListenAndServe to run a new server
	err = srv.ListenAndServe()
	app.errorLog.Fatal(err)
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
