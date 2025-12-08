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
	return nil
}
