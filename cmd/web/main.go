package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"golangs.org/snippetbox/pkg/models/mysql"

	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	snippets *mysql.SnippetModel // add field for access for our handlers
	// components *components
}

func main() {
	// create a new flag cmd
	addr := flag.String("addr", ":80", "Net address HTTP")
	// new cmd flag for setting MySQL connection
	// dsn := flag.String("dsn", "web:ndJMv9zrJw@/snippetbox?parseTime=true", "Name of MySQL data source")
	dsn := flag.String("dsn", "root:Bn7{%14f@/snippetbox?parseTime=true", "Name of MySQL data source")
	// call func for extract flag from cmd
	flag.Parse()

	// log file
	f, err := os.OpenFile("info.log", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	errorLog := log.New(f, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(f, "INFO\t", log.Ldate|log.Ltime)

	// open DB
	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	// set application
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		// initialise instance and add it in depensenses
		snippets: &mysql.SnippetModel{DB: db},
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

// wrapping sql.Open() and return connection pull sql.DB
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
