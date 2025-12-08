package handler

import (
	"device-service/internal/repository"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type HttpHandler struct {
	Db repository.Repo
}

func (h *HttpHandler) Routes() http.Handler {

	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)

	mux.Route("/api/v1/devices", func(r chi.Router) {
		r.Get("/{id}", h.GetDevice)
		r.Post("/", h.AddDevice)
		r.Put("/{id}", h.UpdateDevice)
		r.Delete("/{id}", h.DeleteDevice)

		r.Get("/{id}/sensors/{sensorId}", h.GetSensor)
		r.Post("/{id}/sensors", h.AddSensor)
		r.Put("/{id}/sensors/{sensorId}", h.UpdateSensor)
		r.Delete("/{id}/sensors/{sensorId}", h.DeleteSensor)

	})

	return mux
}
