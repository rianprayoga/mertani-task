package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type HttpHandler struct {
}

func (h *HttpHandler) Routes() http.Handler {

	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)

	mux.Route("/api/v1/devices", func(r chi.Router) {
		r.Get("/{id}", h.GetDevice)
		r.Post("/", h.AddDevice)
		r.Put("/{id}", h.UpdateDevice)
		r.Delete("{id}", h.DeleteDevice)
	})

	return mux
}
