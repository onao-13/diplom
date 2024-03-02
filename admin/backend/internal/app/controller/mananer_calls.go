package controller

import (
	"admin/internal/app/errors"
	"admin/internal/app/handler"
	"admin/internal/app/service"
	"net/http"
)

type ManagerCalls struct {
	service service.ManagerCalls
}

func NewManagerCalls(service service.ManagerCalls) ManagerCalls {
	return ManagerCalls{service: service}
}

func (c ManagerCalls) GetAll(w http.ResponseWriter, r *http.Request) {
	calls, err := c.service.GetAll()
	if err != nil {
		switch err.(type) {
		case errors.ErrNotFound:
			handler.HandleNotFound(w, err.Error())
			return
		default:
			handler.HandlerInternalServerError(w, "Ошибка сервера")
			return
		}
	}

	body := map[string]interface{}{
		"manager_calls": calls,
	}

	handler.HandleOkData(w, body)
}
