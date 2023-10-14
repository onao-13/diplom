package service

import (
	"backend/internal/app/middleware/database"
	"backend/internal/app/payload/response"

	"github.com/sirupsen/logrus"
)

type Home struct {
	db database.Home
	log logrus.Logger	
}

func NewHome(db database.Home, log logrus.Logger) Home {
	return Home{db, log}
}

func (h *Home) GetById(id int64) (response.Home, error) {
	home, err := h.db.GetById(id)
	if err != nil {
		h.log.Errorln("Ошибка получения дома: ", err.Error())
		return response.Home{}, err 
	} 

	return home, nil
}
