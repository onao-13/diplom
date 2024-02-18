package service

import (
	"backend/internal/app/database"
	"backend/internal/app/payload"
	"github.com/sirupsen/logrus"
)

type ManagerCall struct {
	log *logrus.Logger
	db  database.ManagerCall
}

func NewManagerCall(log *logrus.Logger, db database.ManagerCall) ManagerCall {
	return ManagerCall{log: log, db: db}
}

func (m ManagerCall) Save(sendCallRequest payload.SendManagerCall) error {
	if err := m.db.Save(sendCallRequest); err != nil {
		m.log.Errorln("Ошибка сохранения обратного звонка менеджеру: ", err)
		return err
	}
	return nil
}
