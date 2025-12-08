package repository

import "device-service/internal/model"

type Repo interface {
	AddDevice(req model.CreateDeviceReq) (*model.CreateDeviceRes, error)
	UpdateDevice(id string, req model.CreateDeviceReq) (*model.CreateDeviceRes, error)
	GetDevice(deviceId string) (*model.CreateDeviceRes, error)
	DeleteDevice(deviceId string) error

	AddSensor(deviceId string, req model.CreateSensorReq) (*model.CreateSensorRes, error)
	GetSensor(deviceId string, sensorId string) (*model.CreateSensorRes, error)
	UpdateSensor(deviceId string, sensorId string, req model.CreateSensorReq) (*model.CreateSensorRes, error)
	DeleteSensor(deviceId string, sensorId string) error
}
