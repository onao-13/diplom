package controller

import (
	"backend/internal/app/handler"
	"backend/internal/app/service"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Home struct {
	service service.Home
}

func NewHome(s service.Home) Home {
	return Home{s}
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

	home, err := h.service.GetById(id)
	if err != nil {
		handler.HandleNotFound(w, "Дом не найден")
		return
	}

	handler.HandleOkData(w, home)
}
