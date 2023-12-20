package main

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"golang.org/x/net/websocket"
)

func (app *application) routes() http.Handler {
	//create a router mux
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/*", app.Home)
	r.Handle("/countdown", websocket.Handler(app.handleWSOrderbook))

	r.Get("/ping", app.Ping)
	return r
}
