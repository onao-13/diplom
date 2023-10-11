package service

import (
	"admin/internal/app/middleware/database"
	"admin/internal/payload"
)

type City struct {
	db database.City
}

func NewCity(db database.City) City {
	return City{db}
}

func (c *City) Create(city payload.City) error {
	return c.db.Create(city)
}

func (c *City) Preview(id int64) (payload.City, error) {
	return c.db.Preview(id)
}

func (c *City) Update(id int64, city payload.City) error {
	return c.db.Update(id, city)
}

func (c *City) Delete(id int64) error {
	return c.db.Delete(id)
}
