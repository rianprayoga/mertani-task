package repository

import "device-service/internal/model"

type Repo interface {
	AddDevice(req model.CreateDeviceReq) (*model.CreateDeviceRes, error)
	UpdateDevice(id string, req model.CreateDeviceReq) (*model.CreateDeviceRes, error)
	GetDevice(deviceId string) (*model.CreateDeviceRes, error)
	DeleteDevice(deviceId string) error
}
