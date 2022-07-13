package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	"golangs.org/snippetbox/pkg/models"
	"golangs.org/snippetbox/pkg/models/page"
)

// home page handler
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	p := &page.PageData{Title: "Sunney Valley", Keywords: "home", Description: "home page"}

	data := &templateData{
		PageData: p,
	}

	files := []string{
		"./template/html/home/index.html",
		"./template/base.layout.html",
		"./template/header.html",
		"./template/footer.html",
		"./template/includes/header_adress.html",
		"./template/includes/header_email.html",
		"./template/includes/header_logo_img.html",
		"./template/includes/header_logo_name.html",
		"./template/includes/header_phone.html",
		"./template/includes/header_time.html",
		"./template/includes/footer_title_menu_1.html",
		"./template/includes/footer_title_menu_2.html",
		"./template/includes/footer_title_menu_3.html",
		"./pkg/components/hdtopmenu/hdtopmenu.html",
		"./pkg/components/mn_flmenu_top/mn_flmenu_top.html",
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

	// create instance templateData struct included snippet data
	data := &templateData{Snippet: s}

	// slice with paths to html files
	files := []string{
		"./ui/html/show.page.html",
		"./ui/html/base.layout.html",
		"./ui/html/footer.partial.html",
	}

	// parsing template files
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
	}

	// execute template files
	err = ts.Execute(w, data)
	if err != nil {
		app.serverError(w, err)
	}
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
