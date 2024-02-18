package service

import (
	"backend/internal/app/database"
	"backend/internal/app/payload"
	"fmt"
	"github.com/sirupsen/logrus"
)

type City struct {
	log *logrus.Logger
	db  database.City
}

func NewCity(db database.City, log *logrus.Logger) City {
	return City{db: db, log: log}
}

func (c City) GetAll() (cities []*payload.City, err error) {
	if cities, err = c.db.GetAll(); err != nil {
		c.log.Errorln("Ошибка получения списка городов: ", err.Error())
	}

	if len(cities) == 0 {
		return nil, fmt.Errorf("список городов отсутствует")
	}

	return
}
