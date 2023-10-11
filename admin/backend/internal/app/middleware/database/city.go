package database

import (
	"admin/internal/payload"
	"context"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

type City struct {
	ctx context.Context
	pool *pgxpool.Pool
	log logrus.Logger
}

func NewCity(ctx context.Context, pool *pgxpool.Pool, log logrus.Logger) City {
	return City{ctx: ctx, pool: pool, log: log}
}

func (c *City) Create(city payload.City) error {
	sql := `INSERT INTO cities(cityName) VALUES(@city)`
	arg := pgx.NamedArgs{
		"city": city.Name,
	}

	if _, err := c.pool.Exec(c.ctx, sql, arg); err != nil {
		return err
	}

	return nil
}

func (c *City) Preview(id int64) (payload.City, error) {
	var city payload.City

	sql := `SELECT cityName FROM cities WHERE id = $1`

	if err := pgxscan.Select(c.ctx, c.pool, &city, sql, id); err != nil {
		return payload.City{}, err
	}

	return city, nil
}

func (c *City) Update(id int64, city payload.City) error {
	sql := `UPDATE cities SET name=@name WHERE id=@id`

	args := pgx.NamedArgs{
		"name": city.Name,
		"id": id,
	}

	if _, err := c.pool.Exec(c.ctx, sql, args); err != nil {
		return err
	}

	return nil
}

func (c *City) Delete(id int64) error {
	sql := `DELETE FROM cities WHERE id=@id`

	arg := pgx.NamedArgs{
		"id": id,
	}

	if _, err := c.pool.Exec(c.ctx, sql, arg); err != nil {
		return err
	}

	return nil
}
