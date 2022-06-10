package main

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

// home page handler
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	// initialise slice, contained paths to 2 files
	// home.page.html must to be first
	files := []string{
		"C:/Users/Lesik/go/src/snippetbox/ui/html/home.page.html",
		"C:/Users/Lesik/go/src/snippetbox/ui/html/base.layout.html",
		"C:/Users/Lesik/go/src/snippetbox/ui/html/footer.partial.html",
	}

	// use ParseFiles for reading pattern file
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// then use method Exeecute for writing content of pattern
	// in HTTP response body
	err = ts.Execute(w, nil)
	if err != nil {
		app.serverError(w, err)
	}
}

// display note handler

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	// extract value of parameter id from URL
	// and try to convert string to integer using func Atoi
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	// if err, return 404
	if err != nil || id < 0 {
		app.notFound(w)
		return
	}

	// input id value to response
	fmt.Fprintf(w, "Displaying chosen note with ID %d...", id)
}

// create note handler

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	// if not method POST
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	title := "Story about a snale"
	content := "A snale выползла из раковины, \n вытянула ножки, \n и опять подобрала их"
	expires := "7"

	// pass data to method SnippetModel.Insert() and take back created ID
	id, err := app.snippets.Instert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// redirect user to corresponded snippet page
	http.Redirect(w, r, fmt.Sprintf("/snippet?id=%d", id), http.StatusSeeOther)
}
