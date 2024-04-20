package main

import "github.com/go-chi/chi/v5"

func PageRoutes(pageHandler *PageHandler) chi.Router {
	r := chi.NewRouter()
	r.Get("/", pageHandler.ListPages)
	r.Post("/", pageHandler.CreatePage)
	r.Get("/{id}", pageHandler.GetPages)
	r.Put("/{id}", pageHandler.UpdatePage)
	r.Delete("/{id}", pageHandler.DeletePage)
	return r
}
