package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	// initialisation new router
	mux := http.NewServeMux()
	// registaration handlers
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/en/company", app.company)
	mux.HandleFunc("/en/equipment", app.equipment)
	mux.HandleFunc("/en/services", app.services)
	mux.HandleFunc("/en/contacts", app.contacts)
	mux.HandleFunc("/snippet", app.showSnippet)
	mux.HandleFunc("/snippet/create", app.createSnippet)

	// handle http-requests to static files from "/static"
	// fileServer := http.FileServer(neuteredFileSystem{http.Dir("C:/Users/Lesik/go/src/snippetbox/ui/static")})
	fileServer := http.FileServer(http.Dir("C:/Users/Lesik/go/src/snippetbox/template/static"))

	// use for registration handle all requests, begining with "/static/"
	// mux.Handle("/static", http.NotFoundHandler())
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return mux
}
