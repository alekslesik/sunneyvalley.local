package config

import (
	"fmt"
	"net/http"
	"runtime/debug"

)


// helper write error message to errorLog
// then receive to user responce 500
func (app *Application) ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.ErrorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// helper receive status code and description to user
func (app *Application) ClientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}


// wrapper around clientError
// receive to user responce 404
func (app *Application) NotFound(w http.ResponseWriter)  {
	app.ClientError(w, http.StatusNotFound)
}