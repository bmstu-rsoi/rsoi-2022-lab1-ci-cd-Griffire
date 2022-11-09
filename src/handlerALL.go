package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
)

var dbInstance Database

func NewHandler1(db *Database) http.Handler {
	router := chi.NewRouter()
	dbInstance = *db
	router.MethodNotAllowed(methodNotAllowedHandler)
	router.NotFound(notFoundHandler)
	router.Route("/api/v1/persons", persons)
	return router
}
func methodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(204)
	render.Render(w, r, ErrMethodNotAllowed)
}
func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(400)
	render.Render(w, r, ErrNotFound)
}
