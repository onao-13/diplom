package controller

import (
	"backend/internal/app/handler"
	"backend/internal/app/payload"
	"backend/internal/app/service"
	"encoding/json"
	"net/http"
)

type ManagerCall struct {
	service service.ManagerCall
}

func NewManagerCall(service service.ManagerCall) ManagerCall {
	return ManagerCall{service: service}
}

func (m ManagerCall) Send(w http.ResponseWriter, r *http.Request) {
	var managerCall payload.SendManagerCall
	if err := json.NewDecoder(r.Body).Decode(&managerCall); err != nil {
		handler.HandleBadRequest(w, "Ошибка декодирования JSON")
		return
	}

	if err := m.service.Save(managerCall); err != nil {
		handler.HandlerInternalServerError(w, "Ошибка сохранения вызова")
		return
	}

	handler.HandleOkMsg(w, "")
}
