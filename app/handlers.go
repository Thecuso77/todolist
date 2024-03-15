package main

import (
	"net/http"
	"vuego/frontend"
)

func (app *application) indexHandler(w http.ResponseWriter, r *http.Request) {
	rawFle, err := frontend.StaticFiles.ReadFile("dist/index.html")
	if err != nil {
		app.serverError(w, err)
	}
	w.Write(rawFle)
}
