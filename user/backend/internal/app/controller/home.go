package controller

import (
	"backend/internal/app/handler"
	"backend/internal/app/service"
	"github.com/gorilla/mux"
	"net/http"
)

type Home struct {
	service service.Home
}

func NewHome(s service.Home) Home {
	return Home{s}
}

func (h *Home) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		handler.HandleBadRequest(w, "ID пустое")
		return
	}

	home, err := h.service.GetById(id)
	if err != nil {
		handler.HandleNotFound(w, "Дом не найден")
		return
	}

	handler.HandleOkData(w, home)
}
