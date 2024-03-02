package database

import (
	"admin/internal/app/errors"
	"admin/internal/app/payload"
	"context"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type City struct {
	ctx  context.Context
	pool *pgxpool.Pool
}

func NewCity(ctx context.Context, pool *pgxpool.Pool) City {
	return City{ctx: ctx, pool: pool}
}

func (c *City) Create(city payload.City) error {
	sql := `
	INSERT INTO 
	    cities(name) 
	VALUES(@name)`
	arg := pgx.NamedArgs{
		"name": city.Name,
	}

	if _, err := c.pool.Exec(c.ctx, sql, arg); err != nil {
		return err
	}

	return nil
}

func (c *City) Preview(id int64) (payload.City, error) {
	var city payload.City

	sql := `
	SELECT 
	    name 
	FROM 
	    cities 
	WHERE id = $1`

	if err := pgxscan.Get(c.ctx, c.pool, &city, sql, id); err != nil {
		return payload.City{}, err
	}

	return city, nil
}

func (c City) GetAll() (cities []payload.City, err error) {
	sql := `
	SELECT
		id,
		name
	FROM
	    cities`

	rows, err := c.pool.Query(c.ctx, sql)
	if err != nil {
		return
	}

	for rows.Next() {
		var city payload.City
		if err = rows.Scan(&city.Id, &city.Name); err != nil {
			continue
		}
		cities = append(cities, city)
	}

	if len(cities) == 0 {
		return nil, errors.ErrNotFound{}
	}

	return
}

func (c *City) Update(id int64, city payload.City) error {
	sql := `
	UPDATE 
	    cities 
	SET 
	    name=@name 
	WHERE 
	    id=@id`

	args := pgx.NamedArgs{
		"name": city.Name,
		"id":   id,
	}

	if _, err := c.pool.Exec(c.ctx, sql, args); err != nil {
		return err
	}

	return nil
}

func (c *City) Delete(id int64) error {
	sql := `
	DELETE FROM 
	   cities 
   WHERE 
	   id=@id`

	arg := pgx.NamedArgs{
		"id": id,
	}

	if _, err := c.pool.Exec(c.ctx, sql, arg); err != nil {
		return err
	}

	return nil
}
