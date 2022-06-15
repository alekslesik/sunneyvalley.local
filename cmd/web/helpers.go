package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

// helper write error message to errorLog
// then receive to user responce 500
func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// helper receive status code and description to user
func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

// wrapper around clientError
// receive to user responce 404
func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

func (app *application) render(w http.ResponseWriter, r *http.Request, name string, td *templateData)  {
	// extract corresponding pattern set from cache depending on page name
	ts, ok := app.templateCache[name]
	if !ok {
		app.serverError(w, fmt.Errorf("Pattern %s don't exist!", name))
		return
	}

	// render pattern files, received dynamic data from td variable
	err := ts.Execute(w, td)
	if err != nil {
		app.serverError(w, err)
	}
}
