package controller

import (
	"backend/internal/app/handler"
	"backend/internal/app/middleware/service"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Home struct {
	s service.Home
	log logrus.Logger
}

func NewHome(s service.Home, log logrus.Logger) Home {
	return Home{s, log}
}

func (h *Home) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ids, ok := vars["id"]
	if !ok {
		handler.HandleBadRequest(w, "ID пустое")
		return
	}

	id, err := strconv.ParseInt(ids, 10, 64)
	if err != nil {
		handler.HandlerInternalServerError(w, "Ошибка получения ID")
		return
	}

	home, err := h.s.GetById(id)
	if err != nil {
		handler.HandleNotFound(w, "Дом не найден")
		return
	}

	handler.HandleOkData(w, home)
}

func (h *Home) List(w http.ResponseWriter, r *http.Request) {
	
}
