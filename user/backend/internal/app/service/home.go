package service

import (
	"backend/internal/app/database"
	"backend/internal/app/payload"
	"github.com/sirupsen/logrus"
)

type Home struct {
	db  database.Home
	log *logrus.Logger
}

func NewHome(db database.Home, log *logrus.Logger) Home {
	return Home{db, log}
}

func (h Home) GetById(id int64) (home *payload.Home, err error) {
	home, err = h.db.GetById(id)
	if err != nil {
		h.log.Errorln("Ошибка получения дома: ", err.Error())
		return &payload.Home{}, err
	}

	return home, nil
}

func (h Home) GetAll() (homes []payload.Home, err error) {
	//homes, err = h.db.GetAll()
	return
}
