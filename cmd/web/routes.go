package main

import (
	"net/http"

	"github.com/teerkevin23/rundoo/assets"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	mux := httprouter.New()
	mux.NotFound = http.HandlerFunc(app.notFound)

	fileServer := http.FileServer(http.FS(assets.EmbeddedFiles))
	mux.Handler("GET", "/static/*filepath", fileServer)

	mux.HandlerFunc("GET", "/", app.home)
	mux.HandlerFunc("POST", "/createProduct", app.createProduct)
	mux.HandlerFunc("GET", "/getProducts", app.getProducts)


	return app.recoverPanic(app.securityHeaders(mux))
}
