package main

import (
	"bytes"
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

func (app *application) render(w http.ResponseWriter, r *http.Request, name string, td *templateData) {

	// extract pattern depending "name"
	ts, ok := app.templateCache[name]
	if !ok {
		app.serverError(w, fmt.Errorf("pattern %s not exist", name))
		return
	}

	// rendering pattern files passing dynamic data from td variable
	// initialize a new buffer
	buf := new(bytes.Buffer)

	// write template to the buffer, instead straight to http.ResponseWriter
	err := ts.Execute(buf, td)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// write buffer to http.ResponseWriter
	buf.WriteTo(w)
}
