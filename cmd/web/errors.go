package main

import (
	"net/http"
	"runtime/debug"
)

func (app *application) reportError(err error) {
	trace := debug.Stack()

	app.logger.Printf("%s\n%s", err, trace)
}

func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	app.reportError(err)

	message := "The server encountered a problem and could not process your request"
	http.Error(w, message, http.StatusInternalServerError)
}

func (app *application) notFound(w http.ResponseWriter, r *http.Request) {
	message := "The requested resource could not be found"
	http.Error(w, message, http.StatusNotFound)
}

func (app *application) badRequest(w http.ResponseWriter, r *http.Request, err error) {
	http.Error(w, err.Error(), http.StatusBadRequest)
}

func (app *application) conflict(w http.ResponseWriter, r *http.Request, err error) {
	http.Error(w, err.Error(), http.StatusConflict)
}
