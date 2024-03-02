package service

import (
	"admin/internal/app/database"
	"admin/internal/app/errors"
	"admin/internal/app/payload"
	"github.com/sirupsen/logrus"
)

type ManagerCalls struct {
	db  database.ManagerCalls
	log *logrus.Logger
}

func NewManagerCalls(db database.ManagerCalls, log *logrus.Logger) ManagerCalls {
	return ManagerCalls{db: db, log: log}
}

func (c ManagerCalls) GetAll() (calls []payload.ManagerCall, err error) {
	if calls, err = c.db.GetAll(); err != nil || calls == nil {
		c.log.Errorln("Ошибка получения звонков: ", err.Error())
		err = errors.ErrNotFound{}
	}
	return
}
