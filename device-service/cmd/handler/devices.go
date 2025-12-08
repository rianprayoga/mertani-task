package handler

import (
	"database/sql"
	"device-service/internal/model"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (h *HttpHandler) AddDevice(w http.ResponseWriter, r *http.Request) {
	var req model.CreateDeviceReq
	err := h.readJSON(w, r, &req)
	if err != nil {
		h.errorJSON(w, fmt.Errorf("unexpected error"))
		return
	}

	d, err := h.Db.AddDevice(req)
	if err != nil {
		h.errorJSON(w, fmt.Errorf("unexpected error"))
		return
	}
	h.writeJson(w, http.StatusCreated, d)
}

func (h *HttpHandler) UpdateDevice(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	_, err := getDevice(h, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			h.errorJSON(w, fmt.Errorf("device not found"), http.StatusNotFound)
			return
		}
		h.errorJSON(w, fmt.Errorf("unexpected error"))
		return
	}
	var req model.CreateDeviceReq
	err = h.readJSON(w, r, &req)
	if err != nil {
		h.errorJSON(w, fmt.Errorf("unexpected error"))
		return
	}

	res, err := h.Db.UpdateDevice(id, req)
	if err != nil {
		h.errorJSON(w, fmt.Errorf("unexpected error"))
		return
	}

	h.writeJson(w, http.StatusOK, res)
}

func (h *HttpHandler) DeleteDevice(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := h.Db.DeleteDevice(id)
	if err != nil {
		h.errorJSON(w, fmt.Errorf("unexpected error"))
		return
	}
}

func getDevice(h *HttpHandler, id string) (*model.CreateDeviceRes, error) {
	res, err := h.Db.GetDevice(id)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (h *HttpHandler) GetDevice(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	res, err := getDevice(h, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			h.errorJSON(w, fmt.Errorf("device not found"), http.StatusNotFound)
			return
		}
		h.errorJSON(w, fmt.Errorf("unexpected error"))
	}

	h.writeJson(w, http.StatusOK, res)
}
