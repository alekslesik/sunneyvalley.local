package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	"github.com/alekslesik/snippetbox/pkg/models"
	"github.com/alekslesik/snippetbox/pkg/models/form"
	"github.com/alekslesik/snippetbox/pkg/models/page"
	// "github.com/alekslesik/snippetbox/pkg/models/form"
)

// home page handler
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	p := &page.PageData{Title: "Sunney Valley", Keywords: "home", Description: "home page", LeftMenu: false}
	app.render(w, r, "home", &templateData{
		PageData: p,
	})
}

// company page handler
func (app *application) company(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/en/company" {
		app.notFound(w)
		return
	}

	p := &page.PageData{Title: "About company", Keywords: "about", Description: "about", LeftMenu: false}
	app.render(w, r, "company", &templateData{
		PageData: p,
	})
}

// equipment page handler
func (app *application) equipment(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/en/equipment" {
		app.notFound(w)
		return
	}

	p := &page.PageData{Title: "Equipment", Keywords: "equipment", Description: "equipment", LeftMenu: true}

	dirs, err := app.equipments.Dirs()
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	app.render(w, r, "equipment", &templateData{
		PageData:   p,
		Equipments: dirs,
	})
}

// services page handler
func (app *application) services(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/en/services" {
		app.notFound(w)
		return
	}

	p := &page.PageData{Title: "Services", Keywords: "services", Description: "services"}
	app.render(w, r, "services", &templateData{
		PageData: p,
	})
}

// contacts page handler
func (app *application) contacts(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/en/contacts" {
		app.notFound(w)
		return
	}

	p := &page.PageData{Title: "Contacts", Keywords: "contacts", Description: "contacts"}
	app.render(w, r, "contacts", &templateData{
		PageData: p,
	})
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

	s, err := app.equipments.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	// create instance templateData struct included snippet data
	data := &templateData{Equipment: s}

	// slice with paths to html files
	files := []string{
		app.gopath + "/src/github.com/alekslesik/snippetbox/ui/html/show.page.html",
		app.gopath + "/src/github.com/alekslesik/snippetbox/ui/html/base.layout.html",
		app.gopath + "/src/github.com/alekslesik/snippetbox/ui/html/footer.partial.html",
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
	id, err := app.equipments.Instert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// redirect user to corresponded snippet page
	http.Redirect(w, r, fmt.Sprintf("/snippet?id=%d", id), http.StatusSeeOther)
}

// form handler from contact page

func (app *application) mailForm(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/mail-form" {
		app.notFound(w)
		return
	}

	// first parse data from post
	r.ParseForm()

	var formData form.FormData
	formData.SetName(r.PostForm["form_name"][0])
	formData.SetPhone(r.PostForm["form_phone"][0])
	formData.SetMessage(r.PostForm["form_message"][0])

	// http.Redirect(w, r, "/en/contacts", http.StatusMovedPermanently)
	w.Header().Set("Content-Type", "application/json")
	// write the response
	w.Write([]byte("good"))

}
