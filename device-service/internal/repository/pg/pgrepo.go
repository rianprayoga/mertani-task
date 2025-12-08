package pg

import (
	"context"
	"database/sql"
	"device-service/internal/model"
	"time"
)

type PgRepo struct {
	DB *sql.DB
}

const DbTimeout = 3 * time.Second

func (r *PgRepo) AddDevice(req model.CreateDeviceReq) (*model.CreateDeviceRes, error) {

	ctx, cancel := context.WithTimeout(context.Background(), DbTimeout)
	defer cancel()

	var created model.CreateDeviceRes

	err := r.DB.QueryRowContext(
		ctx,
		`
			INSERT INTO devices
			(name, latitude, longitude, created_at, updated_at)
			VALUES ($1, $2, $3, $4, $5)
			RETURNING id, name, latitude, longitude, created_at, updated_at
		`,
		req.Name,
		req.Lat,
		req.Long,
		time.Now().UTC(),
		time.Now().UTC(),
	).Scan(
		&created.Id,
		&created.Name,
		&created.Lat,
		&created.Long,
		&created.CreatedAt,
		&created.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &created, nil

}

func (r *PgRepo) UpdateDevice(id string, req model.CreateDeviceReq) (*model.CreateDeviceRes, error) {

	ctx, cancel := context.WithTimeout(context.Background(), DbTimeout)
	defer cancel()

	var res model.CreateDeviceRes

	err := r.DB.QueryRowContext(
		ctx,
		`UPDATE devices SET
			name = $1,
			latitude = $2,
			longitude = $3,
			updated_at = $4
		WHERE id = $5
			RETURNING id, name, latitude, longitude, created_at, updated_at`,
		req.Name,
		req.Lat,
		req.Long,
		time.Now().UTC(),
		id,
	).Scan(
		&res.Id,
		&res.Name,
		&res.Lat,
		&res.Long,
		&res.CreatedAt,
		&res.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil

}

func (r *PgRepo) GetDevice(deviceId string) (*model.CreateDeviceRes, error) {

	ctx, cancel := context.WithTimeout(context.Background(), DbTimeout)
	defer cancel()

	query := `
			SELECT id, name, latitude, longitude, created_at, updated_at
			FROM devices
			WHERE id = $1
		`

	row := r.DB.QueryRowContext(ctx, query, deviceId)
	var res model.CreateDeviceRes
	err := row.Scan(
		&res.Id,
		&res.Name,
		&res.Lat,
		&res.Long,
		&res.CreatedAt,
		&res.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil

}

func (r *PgRepo) DeleteDevice(deviceId string) error {

	ctx, cancel := context.WithTimeout(context.Background(), DbTimeout)
	defer cancel()

	query := `
			DELETE 
			devices
			WHERE id = $1
		`
	_, err := r.DB.ExecContext(ctx, query, deviceId)
	if err != nil {
		return nil
	}

	return nil
}

func (r *PgRepo) AddSensor(deviceId string, req model.CreateSensorReq) (*model.CreateSensorRes, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DbTimeout)
	defer cancel()

	var created model.CreateSensorRes

	err := r.DB.QueryRowContext(
		ctx,
		`
			INSERT INTO sensors
			(name, device_id, created_at, updated_at)
			VALUES ($1, $2, $3, $4, $5)
			RETURNING id, name, created_at, updated_at
		`,
		req.Name,
		deviceId,
		time.Now().UTC(),
		time.Now().UTC(),
	).Scan(
		&created.Id,
		&created.Name,
		&created.CreatedAt,
		&created.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &created, nil
}

func (r *PgRepo) GetSensor(deviceId string, sensorId string) (*model.CreateSensorRes, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DbTimeout)
	defer cancel()

	var res model.CreateSensorRes

	err := r.DB.QueryRowContext(
		ctx,
		`
		SELECT id, name, created_at, update_at
		FROM sensors 
		WHERE device_id = $1 AND id = $2
		`,
		deviceId,
		sensorId,
	).Scan(
		&res.Id,
		&res.Name,
		&res.CreatedAt,
		&res.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &res, nil

}

func (r *PgRepo) UpdateSensor(deviceId string, sensorId string, req model.CreateSensorReq) (*model.CreateSensorRes, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DbTimeout)
	defer cancel()

	var res model.CreateSensorRes

	err := r.DB.QueryRowContext(
		ctx,
		`UPDATE sensors SET
			name = $1,
			updated_at = $2
		WHERE id = $3 AND device_id = $4
			RETURNING id, name, created_at, updated_at`,
		req.Name,
		time.Now().UTC(),
		sensorId,
		deviceId,
	).Scan(
		&res.Id,
		&res.Name,
		&res.CreatedAt,
		&res.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (r *PgRepo) DeleteSensor(deviceId string, sensorId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), DbTimeout)
	defer cancel()

	query := `
			DELETE 
			sensors
			WHERE id = $1 AND device_id = $2
		`
	_, err := r.DB.ExecContext(ctx, query, sensorId, deviceId)
	if err != nil {
		return err
	}

	return nil
}
