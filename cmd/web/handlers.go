package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"golangs.org/snippetbox/pkg/models"
	"golangs.org/snippetbox/pkg/models/page"
)

// home page handler
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	p := &page.PageData{Title: "Home page", Keywords: "home", Description: "home page"}

<<<<<<< HEAD
	data := &templateData{
		PageData: p,
	}

	files := []string{
		"./template/html/home/index.html",
		"./template/base.layout.html",
		"./template/header.html",
		"./template/footer.html",
	}

	// use ParseFiles for reading pattern file
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// then use method Execute for writing content of pattern
	// in HTTP response body
	err = ts.Execute(w, data)
	if err != nil {
		app.serverError(w, err)
	}
=======
	// use helper render() for displaying pattern
	app.render(w, r, "home.page.html", &templateData{
		Snippets: s,
	})
>>>>>>> master
}

// display snippet handler
func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	// extract value of parameter id from URL
	// and try to convert string to integer using func Atoi
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	// if err, return 404
	if err != nil || id < 0 {
		app.notFound(w)
		return
	}

	s, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	// use helper render() for displaying pattern
	app.render(w, r, "show.page.html", &templateData{
		Snippet: s,
	})
}

// create snippet handler
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
