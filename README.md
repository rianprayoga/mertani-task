
# Task

CRUD for device and sensor.
The relation is one device can have multiple sensors, so it is one to many relationship. 



## API Reference

### Device 

Create device 
```http
  POST /api/v1/devices
```
Get device 
```http
  GET /api/v1/devices/{id}
```
Update device 
```http
  PUT /api/v1/devices/{id}
```
Delete device 
```http
  DELETE /api/v1/devices/{id}
```

### Sensor 

Create Sensor 
```http
  POST /api/v1/devices/{deviceId}/sensors
```
Get sensor 
```http
  GET /api/v1/devices/{deviceId}/sensors/{sensorId}
```
Update sensor
```http
  PUT /api/v1/devices/{deviceId}/sensors/{sensorId}
```
Delete sensor
```http
  DELETE /api/v1/devices/{deviceId}/sensors/{sensorId}
```