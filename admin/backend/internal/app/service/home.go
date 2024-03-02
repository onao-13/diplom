package service

import (
	"admin/internal/app/database"
	"admin/internal/app/errors"
	"admin/internal/app/payload"
	"github.com/sirupsen/logrus"
)

type Home struct {
	log *logrus.Logger
	db  database.Home
}

func NewHome(db database.Home, log *logrus.Logger) Home {
	return Home{db: db, log: log}
}

func (h Home) Create(cityId string, home payload.Home) error {
	if err := h.db.Create(cityId, home); err != nil {
		h.log.Errorln("Ошибка создания дома: ", err.Error())
		return err
	}
	return nil
}

func (h Home) GetById(cityId, homeId string) (home *payload.Home, err error) {
	home, err = h.db.GetById(cityId, homeId)
	if err != nil {
		h.log.Errorln("Ошибка получение дома по ID: ", err.Error())
	}
	if home == nil {
		return nil, errors.ErrNotFound{}
	}
	return
}

func (h Home) GetByCityId(cityId string) (homes []*payload.Home, err error) {
	homes, err = h.db.GetByCityId(cityId)
	if err != nil {
		h.log.Errorln("Ошибка получения списка домов по ID городу: ", err.Error())
	}

	if len(homes) == 0 {
		return nil, errors.ErrNotFound{}
	}
	return
}

func (h Home) Update(homeId string, home payload.HomeUpdate) error {
	if err := h.db.Update(homeId, home); err != nil {
		h.log.Errorln("Ошибка обновления дома: ", err.Error())
		return err
	}
	return nil
}

func (h Home) Delete(cityId, homeId string) error {
	if err := h.db.Delete(cityId, homeId); err != nil {
		h.log.Errorln("Ошибка удаления дома: ", err.Error())
		return err
	}
	return nil
}
