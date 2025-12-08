package handler

import (
	"database/sql"
	"device-service/internal/model"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (h *HttpHandler) AddSensor(w http.ResponseWriter, r *http.Request) {
	deviceId := chi.URLParam(r, "id")

	_, err := h.GetDeviceById(deviceId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			h.errorJSON(w, fmt.Errorf("device not found"), http.StatusNotFound)
			return
		}
		h.errorJSON(w, fmt.Errorf("unexpected error"))
		return
	}

	var req model.CreateSensorReq
	err = h.readJSON(w, r, &req)
	if err != nil {
		h.errorJSON(w, fmt.Errorf("unexpected error"))
		return
	}

	res, err := h.Db.AddSensor(deviceId, req)
	if err != nil {
		h.errorJSON(w, fmt.Errorf("unexpected error"))
		return
	}

	h.writeJson(w, http.StatusCreated, res)

}

func (h *HttpHandler) GetSensor(w http.ResponseWriter, r *http.Request) {
	deviceId := chi.URLParam(r, "id")

	_, err := h.GetDeviceById(deviceId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			h.errorJSON(w, fmt.Errorf("device not found"), http.StatusNotFound)
			return
		}
		h.errorJSON(w, fmt.Errorf("unexpected error"))
		return
	}

	sensorId := chi.URLParam(r, "sensorId")

	res, err := h.Db.GetSensor(deviceId, sensorId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			h.errorJSON(w, fmt.Errorf("sensor not found"), http.StatusNotFound)
			return
		}
		h.errorJSON(w, fmt.Errorf("unexpected error"))
		return
	}

	h.writeJson(w, http.StatusOK, res)
}

func (h *HttpHandler) UpdateSensor(w http.ResponseWriter, r *http.Request) {
	deviceId := chi.URLParam(r, "id")

	_, err := h.GetDeviceById(deviceId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			h.errorJSON(w, fmt.Errorf("device not found"), http.StatusNotFound)
			return
		}
		h.errorJSON(w, fmt.Errorf("unexpected error"))
		return
	}

	sensorId := chi.URLParam(r, "sensorId")

	_, err = h.Db.GetSensor(deviceId, sensorId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			h.errorJSON(w, fmt.Errorf("sensor not found"), http.StatusNotFound)
			return
		}
		h.errorJSON(w, fmt.Errorf("unexpected error"))
		return
	}
}

func (h *HttpHandler) DeleteSensor(w http.ResponseWriter, r *http.Request) {
	deviceId := chi.URLParam(r, "id")

	_, err := h.GetDeviceById(deviceId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			h.errorJSON(w, fmt.Errorf("device not found"), http.StatusNotFound)
			return
		}
		h.errorJSON(w, fmt.Errorf("unexpected error"))
		return
	}

	sensorId := chi.URLParam(r, "sensorId")

	err = h.Db.DeleteSensor(deviceId, sensorId)
	if err != nil {
		h.errorJSON(w, fmt.Errorf("unexpected error"))
		return
	}

	h.writeJson(w, http.StatusOK, &JSONResponse{
		Message: "sensor deleted",
	})
}
