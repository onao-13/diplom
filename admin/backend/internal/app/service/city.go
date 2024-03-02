package service

import (
	"admin/internal/app/database"
	"admin/internal/app/payload"
	"github.com/sirupsen/logrus"
)

type City struct {
	log *logrus.Logger
	db  database.City
}

func NewCity(log *logrus.Logger, db database.City) City {
	return City{
		log: log,
		db:  db,
	}
}

func (c *City) Create(city payload.City) error {
	if err := c.db.Create(city); err != nil {
		c.log.Errorln("Ошибка создания города: ", err.Error())
		return err
	}
	return nil
}

func (c City) GetAll() (cities []payload.City, err error) {
	cities, err = c.db.GetAll()
	if err != nil {
		c.log.Errorln("Ошибка получения списка городов: ", err.Error())
		return
	}
	return
}

func (c *City) Preview(id int64) (city payload.City, err error) {
	city, err = c.db.Preview(id)
	if err != nil {
		c.log.Errorln("Ошибка получения статьи: ", err.Error())
		return
	}
	return
}

func (c *City) Update(id int64, city payload.City) error {
	if err := c.db.Update(id, city); err != nil {
		c.log.Errorln("Ошибка обновления города: ", err.Error())
		return err
	}
	return nil
}

func (c *City) Delete(id int64) error {
	if err := c.db.Delete(id); err != nil {
		c.log.Errorln("Ошибка удаления города: ", err.Error())
		return err
	}
	return nil
}
